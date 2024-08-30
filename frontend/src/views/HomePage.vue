<template>
    <div class="home">
        <h1>發起活動</h1>
        <form @submit.prevent="submitForm">
            <div class="form-group">
                <label for="title">活動標題:</label>
                <input id="title" v-model="form.title" type="text" required>
            </div>
            <div class="form-group">
                <label for="description">活動內容:</label>
                <textarea id="description" v-model="form.description" required></textarea>
            </div>
            <div class="form-group">
                <label for="location">活動地點:</label>
                <input id="location" v-model="form.location" type="text" required>
                <!-- 只有在地點有文字时才显示地圖 -->
                <div v-if="form.location" class="map-container">
                    <iframe width="100%" height="300" frameborder="0" scrolling="no" marginheight="0" marginwidth="0"
                        :src="'https://maps.google.com/maps?width=100%25&amp;height=300&amp;hl=zh-TW&amp;q=' + encodeURIComponent(form.location) + '+(My%20Business%20Name)&amp;t=&amp;z=14&amp;ie=UTF8&amp;iwloc=B&amp;output=embed&disableDefaultUI=true&zoomControl=false&mapTypeControl=false&scaleControl=false&streetViewControl=false'"
                        allowfullscreen aria-hidden="false" tabindex="0"></iframe>
                </div>
            </div>
            <div class="form-group">
                <label for="time">活動時段:</label>
                <input id="time" v-model="form.time" type="datetime-local" required>
            </div>
            <div class="form-group">
                <label for="image">活動圖片:</label>
                <input id="image" @change="onFileChange" type="file" required>
            </div>
            <button type="submit" class="submit-btn">提交活動</button>
        </form>
    </div>
</template>

<script>
export default {
    name: 'HomePage',
    data() {
        return {
            form: {
                title: '',
                description: '',
                location: '',
                time: '',
                image: null
            }
        };
    },
    methods: {
        onFileChange(e) {
            this.form.image = e.target.files[0];
        },
        submitForm() {
            let formData = new FormData();
            formData.append('image', this.form.image);

            fetch('http://localhost:8081/api/upload/image', {
                method: 'POST',
                body: formData,
            })
                .then(response => {
                    if (!response.ok) {
                        throw new Error('Network response was not ok');
                    }
                    return response.json();
                })
                .then(data => {
                    if (data.path) {
                        console.log("Image uploaded, URL:", data.path);
                        this.submitEventDetails(data.path);
                    } else {
                        throw new Error('Failed to upload image. No image URL provided.');
                    }
                })
                .catch(error => {
                    console.error('Error uploading image:', error);
                    alert('圖片上傳過程中出現錯誤！');
                });
        },
        submitEventDetails(imageUrl) {
            let dateObject = new Date(this.form.time);
            let formattedTime = dateObject.toISOString();

            let detailsData = {
                title: this.form.title,
                description: this.form.description,
                location: this.form.location,
                time: formattedTime,
                image_url: imageUrl
            };

            fetch('http://localhost:8081/api/activitys/', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(detailsData),
            })
                .then(response => {
                    if (!response.ok) {
                        throw new Error('Network response was not ok: ' + response.statusText);
                    }
                    return response.json();
                })
                .then(data => {
                    console.log("Activity submitted with image URL:", data);
                    alert('活動成功提交！');
                    this.resetForm();
                })
                .catch(error => {
                    console.error('Error submitting activity:', error);
                    alert('活動提交過程中出現錯誤！');
                });
        },
        resetForm() {
            this.form.title = '';
            this.form.description = '';
            this.form.location = '';
            this.form.time = '';
            this.form.image = null;
        }
    }
}
</script>

<style scoped>
.home {
    max-width: 700px;
    margin: 40px auto;
    padding: 30px;
    background: #ffffff;
    border-radius: 12px;
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.1);
}

h1 {
    text-align: center;
    font-size: 2rem;
    color: #333;
    margin-bottom: 30px;
}

.form-group {
    margin-bottom: 20px;
}

label {
    display: block;
    margin-bottom: 8px;
    font-weight: bold;
    color: #333;
}

input[type="text"],
input[type="datetime-local"],
textarea,
input[type="file"] {
    width: 100%;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 6px;
    font-size: 1rem;
    color: #333;
    box-sizing: border-box;
    background: #fafafa;
    transition: border-color 0.3s ease;
}

input[type="text"]:focus,
input[type="datetime-local"]:focus,
textarea:focus,
input[type="file"]:focus {
    border-color: #0056b3;
    outline: none;
}

textarea {
    height: 120px;
}

.map-container {
    width: 100%;
    margin-top: 15px;
    border-radius: 8px;
    overflow: hidden;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.05);
}

.submit-btn {
    display: block;
    width: 100%;
    padding: 15px;
    background-color: #0056b3;
    color: white;
    border: none;
    border-radius: 6px;
    font-size: 1.25rem;
    font-weight: bold;
    cursor: pointer;
    transition: background-color 0.3s ease, transform 0.3s ease;
}

.submit-btn:hover {
    background-color: #004494;
    transform: translateY(-2px);
}

.submit-btn:active {
    background-color: #003366;
    transform: translateY(0);
}
</style>
