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
                    if (data.path) { // 修改此處以使用實際的返回鍵 'path'
                        console.log("Image uploaded, URL:", data.path);
                        this.submitEventDetails(data.path); // 使用圖片存儲路徑作為參數呼叫 submitEventDetails
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
            let dateObject = new Date(this.form.time);  // Convert string to Date object
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
    max-width: 600px;
    margin: 0 auto;
    padding: 20px;
    background: #f9f9f9;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

h1 {
    color: #333;
}

.form-group {
    margin-bottom: 15px;
}

label {
    display: block;
    margin-bottom: 5px;
    color: #333;
}

input[type="text"],
input[type="datetime-local"],
textarea,
input[type="file"] {
    width: 100%;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 4px;
}

input[type="file"] {
    padding: 5px;
}

textarea {
    height: 100px;
}

.submit-btn {
    padding: 10px 20px;
    background-color: #0056b3;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s;
}

.submit-btn:hover {
    background-color: #004494;
}
</style>
