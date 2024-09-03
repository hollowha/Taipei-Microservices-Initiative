<template>
  <div class="vote-container">
    <div v-for="option in options" :key="option.id" class="option">
      <button @click="vote(option.id)">{{ option.label }}</button>
      <div v-if="showResults">
        <div class="progress">
          <div class="progress-bar" :style="{ width: calculatePercentage(option.votes) + '%' }">
            {{ option.votes }} 票 ({{ calculatePercentage(option.votes) }}%)
          </div>
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
      this.$emit('vote', optionId);  // 發送投票事件給父元件
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
}
.option {
  margin-bottom: 10px;
}
.progress {
  background-color: #e0e0e0;
  border-radius: 2px;
  overflow: hidden;
}
.progress-bar {
  background-color: #3b82f6;
  height: 20px;
  line-height: 20px;
  text-align: center;
  color: white;
}
</style>
