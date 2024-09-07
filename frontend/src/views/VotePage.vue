<template>
  <div class="vote-page">
    <h2 class="page-title">活動投票</h2>
    <div v-if="activities.length === 0" class="empty-message">目前沒有可供投票的活動。</div>
    <div v-for="activity in activities" :key="activity.id" class="activity-container">
      <h3 class="activity-title">{{ activity.name }}</h3>
      <VoteOptions
        :options="activity.options"
        :showResults="true"
        @vote="handleVote"
      />
    </div>
  </div>
</template>

<script>
import VoteOptions from '@/components/VoteOptions.vue';
import axios from 'axios';

export default {
  data() {
    return {
      activities: [], // 保存從後端獲取的活動及其選項
    };
  },
  components: {
    VoteOptions
  },
  methods: {
    // 從後端獲取活動及其選項
    fetchActivities() {
      axios.get('http://localhost:8081/api/vote/activities')  // 後端 API 路徑
        .then(response => {
          this.activities = response.data;
        })
        .catch(error => {
          console.error("無法獲取活動數據", error);
        });
    },

    // 提交投票
    handleVote(optionId) {
      axios.post('http://localhost:8081/api/vote/submit', { option_id: optionId })  // 提交投票後重新獲取活動數據
        .then(() => {
          this.fetchActivities();
        })
        .catch(error => {
          console.error("投票失敗", error);
        });
    }
  },
  created() {
    this.fetchActivities(); // 頁面加載時獲取活動及其選項
  }
};
</script>

<style scoped>
.vote-page {
  max-width: 800px;
  margin: 0 auto;
  padding: 30px; /* 增加內邊距 */
  background-color: #f8f9fa;
  border-radius: 20px; /* 圓角加大 */
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15); /* 加強陰影 */
}

.page-title {
  text-align: center;
  color: #343a40;
  font-size: 32px; /* 加大字體 */
  font-weight: bold; /* 字體加粗 */
  margin-bottom: 40px;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1); /* 增加字體陰影 */
}

.empty-message {
  text-align: center;
  color: #6c757d;
  font-size: 20px;
  padding: 30px;
}

.activity-container {
  margin-bottom: 40px;
}

.activity-title {
  font-size: 24px; /* 加大字體 */
  font-weight: bold;
  color: #007bff;
  margin-bottom: 20px;
  text-align: center;
  border-bottom: 2px solid #007bff; /* 增加底線 */
  padding-bottom: 10px; /* 增加底部距離 */
}
</style>
