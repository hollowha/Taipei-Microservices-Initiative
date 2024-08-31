<template>
    <div class="app">
        <h1>GPS 位置和麥克風錄音</h1>

        <div>
            <button @click="getLocation">獲取位置</button>
            <p v-if="location">
                當前位置：緯度: {{ location.latitude }}, 經度: {{ location.longitude }}
            </p>
        </div>

        <div>
            <button @click="startRecording" :disabled="isRecording">開始錄音</button>
            <button @click="stopRecording" :disabled="!isRecording">停止錄音</button>
            <p v-if="audioUrl">
                <audio :src="audioUrl" controls></audio>
            </p>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            location: null,
            isRecording: false,
            mediaRecorder: null,
            audioChunks: [],
            audioUrl: null,
        };
    },
    methods: {
        getLocation() {
            if (navigator.geolocation) {
                navigator.geolocation.getCurrentPosition(
                    (position) => {
                        this.location = {
                            latitude: position.coords.latitude,
                            longitude: position.coords.longitude,
                        };
                    },
                    (error) => {
                        console.error("Error getting location:", error);
                        alert("無法獲取位置，請檢查是否已允許位置權限。");
                    }
                );
            } else {
                alert("你的裝置不支援地理位置功能");
            }
        },
        async startRecording() {
            if (navigator.mediaDevices && navigator.mediaDevices.getUserMedia) {
                try {
                    const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
                    this.mediaRecorder = new MediaRecorder(stream);
                    this.mediaRecorder.start();
                    this.isRecording = true;
                    this.audioChunks = [];

                    this.mediaRecorder.ondataavailable = (e) => {
                        this.audioChunks.push(e.data);
                    };

                    this.mediaRecorder.onstop = () => {
                        const blob = new Blob(this.audioChunks, { type: "audio/wav" });
                        this.audioUrl = URL.createObjectURL(blob);
                    };
                } catch (error) {
                    console.error("Error accessing microphone:", error);
                    alert("無法獲取麥克風權限，請檢查是否已允許麥克風權限。");
                }
            } else {
                alert("你的裝置不支援錄音功能");
            }
        },
        stopRecording() {
            if (this.mediaRecorder && this.isRecording) {
                this.mediaRecorder.stop();
                this.isRecording = false;
            }
        },
    },
};
</script>

<style scoped>
.app {
    max-width: 600px;
    margin: 20px auto;
    padding: 20px;
    background: #f9f9f9;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

h1 {
    color: #333;
    text-align: center;
    margin-bottom: 20px;
}

button {
    padding: 10px 20px;
    margin: 10px;
    background-color: #0056b3;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s ease;
}

button:hover {
    background-color: #004494;
}

button:disabled {
    background-color: #ccc;
    cursor: not-allowed;
}

audio {
    display: block;
    margin-top: 20px;
}
</style>