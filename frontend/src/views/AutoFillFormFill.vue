<template>
    <div>
      <h1>{{ formType }}</h1>
      <p>顯示 {{ formType }} 對應的資料</p>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        formTitle: '',
        formType: this.$route.params.formType
      };
    },
    created() {
      this.loadFormData();
    },
    watch: {
      '$route.params.formType': 'loadFormData'
    },
    methods: {
      loadFormData() {
        this.formType = this.$route.params.formType;
        
        
      }
    }
  }
  </script>
  