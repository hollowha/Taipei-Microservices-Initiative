<template>
    <div class="report-submit">
        <img
            src="@/assets/bloodtrail.png"
            alt="Report"
            class="report-logo"
            @click="hadleClick"
        />
        <h1>提交建議或問題</h1>
        <form @submit.prevent="submitForm">
            <div class="form-group">
                <label for="title">標題:</label>
                <input id="title" v-model="form.title" type="text" required/>
            </div>
            <div class="form-group">
                <label for="type">類型:</label>
                <select class="report-type" id="type" v-model="form.type" required>
                    <option value="">請選擇</option>
                    <option value="問題">問題</option>
                    <option value="建議">建議</option>
                </select>
            </div>
            <div class="form-group">
                <label for="description">描述:</label>
                <textarea id="description" v-model="form.description" required></textarea>
            </div>
            <button type="submit" class="submit-btn">提交</button>
        </form>
    </div>
</template>

<script>
export default {
    name: 'ReportPage',
    data() {
        return {
            form: {
                title: '',
                type: '',
                description: '',
                time: '',
            },
            logoClickCount: 0,
        };
    },
    methods: {
        hadleClick() {
            this.logoClickCount++;
            if (this.logoClickCount >= 5) {
                this.logoClickCount = 0;
                this.$router.push('/report-manage');
            }
        },
        submitForm() {
            let currnetTime = new Date();
            let formData = new FormData();
            formData = {
                title: this.form.title,
                type: this.form.type,
                description: this.form.description, 
                time: currnetTime.toISOString()
            };
            fetch('http://localhost:8081/api/report/', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData),
            })
                .then(response => {
                    if (!response.ok) {
                        throw new Error('Network response was not ok');
                    }
                    return response.json();
                })
                .then(data => {
                    console.log('Report submitted:', data);
                    alert('感謝您的回報！');
                    this.resetForm();
                })
                .catch(error => {
                    console.error('Error submitting report:', error);
                    alert('回報過程中出現錯誤！');
                });
        },
        resetForm() {
            this.form.title = '';
            this.form.type = '';
            this.form.description = '';
        },
    },
}
</script>

<style scoped>
.report-logo {
    display: block;
    margin: 0 auto;
    width: 100px;
    height: 100px;
    border-radius: 50%;
    cursor: default;
}
.report-type {
    width: 100%;
    padding: 10px;
    font-size: 1rem;
    border: 1px solid #ccc;
    border-radius: 6px;
}

.report-submit {
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
textarea {
    width: 100%;
    padding: 10px;
    font-size: 1rem;
    border: 1px solid #ccc;
    border-radius: 6px;
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
textarea {
    height: 120px;
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