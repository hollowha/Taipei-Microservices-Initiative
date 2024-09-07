<template>
    <div class="auto-fill-form">
        <div class="button-container">
            <button class="button" @click="goToForm('Identity Card Application Form')">Identity Card Application Form</button>
            <button class="button" @click="goToForm('Address Change Application Form')">Address Change Application Form</button>
            <button class="button" @click="goToForm('simpleform')">simpleform</button>
        </div>
    </div>
</template>
<script>
export default {
  methods: {
    goToForm(formType) {
      this.$router.push(`/auto-fill-form/${formType}`);
    }
  }
}
</script>

<style scoped>
.auto-fill-form {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh; /* Adjust as needed */
}

.button-container {
    display: flex;
    flex-direction: column;
    gap: 10px; /* Adjust spacing between buttons as needed */
}

.button {
    padding: 10px 20px;
    font-size: 16px;
}
</style>
