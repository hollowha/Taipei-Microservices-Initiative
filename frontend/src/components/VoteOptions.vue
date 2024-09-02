<template>
    <div class="vote-container">
      <h3>{{ title }}</h3>
      <div v-for="(option, index) in options" :key="index" class="option">
        <button @click="vote(index)">{{ option.label }}</button>
        <div v-if="showResults">
          <div class="progress">
            <div class="progress-bar" :style="{ width: option.percentage + '%' }">
              {{ option.percentage }}%
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      title: String,
      options: Array,
      showResults: Boolean
    },
    methods: {
      vote(index) {
        this.$emit('vote', index);
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
  