<template>
    <div class="create-activity">
      <h2>新增活動</h2>
      <form @submit.prevent="submitActivity" class="activity-form">
        <div class="form-group">
          <label for="activity-name">活動名稱：</label>
          <input type="text" id="activity-name" v-model="activityName" placeholder="請輸入活動名稱" required />
        </div>
        <div class="form-group">
          <label>選項：</label>
          <div v-for="(option, index) in options" :key="index" class="option-group">
            <input type="text" v-model="options[index].label" placeholder="輸入選項名稱" required />
            <button type="button" @click="removeOption(index)" class="btn-remove">移除</button>
          </div>
          <button type="button" @click="addOption" class="btn-add-option">新增選項</button>
        </div>
        <button type="submit" class="btn-submit">提交活動</button>
      </form>
      <p v-if="successMessage" class="success">{{ successMessage }}</p>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        activityName: '', // 活動名稱
        options: [{ label: '' }], // 活動的選項
        successMessage: '', // 成功訊息
      };
    },
    methods: {
      // 新增選項
      addOption() {
        this.options.push({ label: '' });
      },
      // 移除選項
      removeOption(index) {
        this.options.splice(index, 1);
      },
      // 提交活動
      submitActivity() {
        const payload = {
          name: this.activityName,
          options: this.options.map(option => ({ label: option.label }))
        };
  
        axios.post('http://localhost:8081/api/vote/create', payload)
            .then(() => {
                this.successMessage = '活動已成功新增！';
                this.activityName = '';
                this.options = [{ label: '' }];
            })
        .catch(error => {
            console.error("無法新增活動", error);
        });
      }
    }
  };
  </script>
  
  <style scoped>
  .create-activity {
    max-width: 600px;
    margin: 0 auto;
    padding: 30px;
    background-color: #f8f9fa;
    border-radius: 10px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  }
  
  h2 {
    text-align: center;
    margin-bottom: 20px;
    color: #343a40;
  }
  
  .activity-form {
    display: flex;
    flex-direction: column;
  }
  
  .form-group {
    margin-bottom: 20px;
  }
  
  label {
    font-weight: 600;
    margin-bottom: 10px;
    color: #495057;
  }
  
  input[type="text"] {
    width: 100%;
    padding: 10px;
    border: 1px solid #ced4da;
    border-radius: 5px;
    font-size: 16px;
    color: #495057;
  }
  
  input[type="text"]:focus {
    border-color: #80bdff;
    outline: none;
  }
  
  .option-group {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 10px;
  }
  
  .btn-remove {
    background-color: #dc3545;
    color: white;
    border: none;
    padding: 8px 12px;
    border-radius: 5px;
    cursor: pointer;
    font-size: 14px;
  }
  
  .btn-remove:hover {
    background-color: #c82333;
  }
  
  .btn-add-option {
    background-color: #007bff;
    color: white;
    padding: 10px 15px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    font-size: 16px;
  }
  
  .btn-add-option:hover {
    background-color: #0056b3;
  }
  
  .btn-submit {
    background-color: #28a745;
    color: white;
    padding: 10px 15px;
    border: none;
    border-radius: 5px;
    font-size: 16px;
    font-weight: 700;
    cursor: pointer;
  }
  
  .btn-submit:hover {
    background-color: #218838;
  }
  
  .success {
    text-align: center;
    color: #28a745;
    font-size: 16px;
    margin-top: 20px;
  }
  </style>
  