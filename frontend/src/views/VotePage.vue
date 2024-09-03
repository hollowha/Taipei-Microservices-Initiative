<template>
  <div class="vote-page">
    <div v-for="activity in activities" :key="activity.id" class="activity">
      <h2>{{ activity.name }}</h2>
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
      axios.get('http://localhost:8081/api/vote/activities')  // 使用後端提供的 API
        .then(response => {
          this.activities = response.data;
        })
        .catch(error => {
          console.error("無法獲取活動數據", error);
        });
    },

    // 提交投票
    handleVote(optionId) {
      axios.post('http://localhost:8081/api/vote/submit', { option_id: optionId })  // 使用後端提供的 API
        .then(() => {
          this.fetchActivities(); // 提交投票後重新獲取活動數據，以顯示更新後的投票數
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
  padding: 20px;
}
.activity {
  margin-bottom: 20px;
}
</style>
