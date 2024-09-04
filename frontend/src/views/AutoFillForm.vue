<template>
    <div class="auto-fill-form">
        <div class="button-container">
            <button class="button" @click="button1Clicked">Button 1</button>
            <button class="button" @click="button2Clicked">Button 2</button>
            <button class="button" @click="button3Clicked">Button 3</button>
        </div>
    </div>
</template>

<script>
export default {
    methods: {
        button1Clicked() {
            console.log("Button 1 is pressed");
            window.open('http://localhost:8081/api/autofillform/1', '_blank');
        },
        button2Clicked() {
            console.log("Button 2 is pressed");
        },
        button3Clicked() {
            console.log("Button 3 is pressed");
        },
    },
};
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