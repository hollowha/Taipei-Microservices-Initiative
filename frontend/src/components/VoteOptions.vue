<template>
  <div class="vote-container">
    <div v-for="option in options" :key="option.id" class="option-container">
      <button @click="vote(option.id)" class="vote-btn">{{ option.label }}</button>
      <div v-if="showResults" class="progress-bar-container">
        <div class="progress-bar" :style="{ width: calculatePercentage(option.votes) + '%' }">
          {{ option.votes }} 票 ({{ calculatePercentage(option.votes) }}%)
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    options: Array,
    showResults: Boolean
  },
  methods: {
    vote(optionId) {
      this.$emit('vote', optionId); // 當用戶點擊投票按鈕時，觸發 'vote' 事件並傳遞選項的 ID
    },
    calculatePercentage(votes) {
      const totalVotes = this.options.reduce((total, option) => total + option.votes, 0);
      return totalVotes ? ((votes / totalVotes) * 100).toFixed(2) : 0;
    }
  }
}
</script>

<style scoped>
.vote-container {
  margin: 20px;
  background-color: #ffffff;
  border: 1px solid #e9ecef;
  padding: 15px;
  border-radius: 15px; /* 圓角加大 */
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1); /* 增加陰影 */
}

.option-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 15px;
}

.vote-btn {
  background: linear-gradient(135deg, #007bff 0%, #0056b3 100%); /* 漸層背景 */
  color: white;
  border: none;
  padding: 12px 24px; /* 加大按鈕 */
  border-radius: 8px; /* 圓角加大 */
  cursor: pointer;
  font-size: 18px; /* 加大字體 */
  transition: transform 0.3s ease, background-color 0.3s ease; /* 增加動畫 */
}

.vote-btn:hover {
  background-color: #003a70;
  transform: translateY(-2px); /* 提升按鈕的互動感 */
}

.progress-bar-container {
  width: 100%;
  background-color: #e9ecef;
  border-radius: 8px;
  margin-top: 10px;
}

.progress-bar {
  background: linear-gradient(90deg, #28a745, #218838); /* 漸層進度條 */
  height: 25px; /* 增加高度 */
  line-height: 25px;
  text-align: center;
  color: white;
  border-radius: 8px;
  transition: width 0.3s ease;
}
</style>
