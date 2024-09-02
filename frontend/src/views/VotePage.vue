<template>
    <div class="vote-page">
      <h1>{{ voteTitle }}</h1>
      <VoteOptions
        :title="voteTitle"
        :options="voteOptions"
        :showResults="showResults"
        @vote="handleVote"
      />
    </div>
  </template>
  
  <script>
import VoteOptions from '@/components/VoteOptions.vue';

  
  export default {
    data() {
      return {
        voteTitle: '請選擇您支持的選項',
        voteOptions: [
          { label: '選項1', votes: 0, percentage: 0 },
          { label: '選項2', votes: 0, percentage: 0 },
          { label: '選項3', votes: 0, percentage: 0 }
        ],
        showResults: false
      };
    },
    components: {
      VoteOptions
    },
    methods: {
      handleVote(index) {
        this.voteOptions[index].votes++;
        this.updatePercentages();
        this.showResults = true;
      },
      updatePercentages() {
        const totalVotes = this.voteOptions.reduce((sum, option) => sum + option.votes, 0);
        this.voteOptions.forEach(option => {
          option.percentage = totalVotes ? ((option.votes / totalVotes) * 100).toFixed(2) : 0;
        });
      }
    }
  };
  </script>
  
  <style scoped>
  .vote-page {
    padding: 20px;
  }
  </style>
  