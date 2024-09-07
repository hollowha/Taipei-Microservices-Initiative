<template>
    <div class="app">
        <h1>GPS 位置、相機錄影和麥克風錄音</h1>

        <!-- 顯示錯誤訊息 -->
        <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>

        <div>
            <button @click="getLocation">獲取位置</button>
            <p v-if="location">
                當前位置：緯度: {{ location.latitude }}, 經度: {{ location.longitude }}
            </p>
        </div>

        <!-- 錄影部分 -->
        <div>
            <h2>錄影部分</h2>
            <button @click="startVideoRecording" :disabled="isRecording">開始錄影</button>
            <button @click="stopVideoRecording" :disabled="!isRecording">停止錄影</button>
            <p v-if="videoUrl">
                <video :src="videoUrl" controls></video>
            </p>
        </div>

        <!-- 錄音部分 -->
        <div>
            <h2>錄音部分</h2>
            <button @click="startAudioRecording" :disabled="isAudioRecording">開始錄音</button>
            <button @click="stopAudioRecording" :disabled="!isAudioRecording">停止錄音</button>
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
            isAudioRecording: false,
            mediaRecorder: null,
            videoChunks: [],
            audioChunks: [],
            videoUrl: null,
            audioUrl: null,
            errorMessage: '', // 顯示錯誤訊息
        };
    },
    methods: {
        getLocation() {
            this.errorMessage = ''; // 清除之前的錯誤訊息
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
                        this.errorMessage = "無法獲取位置，請檢查是否已允許位置權限。";
                    }
                );
            } else {
                this.errorMessage = "你的裝置不支援地理位置功能";
            }
        },

        // 開始錄影
        async startVideoRecording() {
            this.errorMessage = ''; // 清除之前的錯誤訊息
            if (navigator.mediaDevices && navigator.mediaDevices.getUserMedia) {
                try {
                    const stream = await navigator.mediaDevices.getUserMedia({ video: true });
                    this.mediaRecorder = new MediaRecorder(stream);
                    this.mediaRecorder.start();
                    this.isRecording = true;
                    this.videoChunks = [];

                    this.mediaRecorder.ondataavailable = (e) => {
                        this.videoChunks.push(e.data);
                    };

                    this.mediaRecorder.onstop = () => {
                        const blob = new Blob(this.videoChunks, { type: "video/webm" });
                        this.videoUrl = URL.createObjectURL(blob);
                    };
                } catch (error) {
                    console.error("Error accessing camera:", error);

                    if (error.name === 'NotReadableError') {
                        this.errorMessage = "無法啟動相機，可能有其他應用正在使用相機。" + error;
                    } else if (error.name === 'NotAllowedError' || error.name === 'SecurityError') {
                        this.errorMessage = "無法獲取相機權限，請檢查是否已允許相機權限。" + error;
                    } else {
                        this.errorMessage = "相機錄影失敗，請檢查您的裝置或瀏覽器設置。";
                    }
                }
            } else {
                this.errorMessage = "你的裝置不支援相機錄影功能";
            }
        },

        // 停止錄影
        stopVideoRecording() {
            if (this.mediaRecorder && this.isRecording) {
                this.mediaRecorder.stop();
                this.isRecording = false;
            }
        },

        // 開始錄音
        async startAudioRecording() {
            this.errorMessage = ''; // 清除之前的錯誤訊息
            if (navigator.mediaDevices && navigator.mediaDevices.getUserMedia) {
                try {
                    const stream = await navigator.mediaDevices.getUserMedia({
                        audio: {
                            echoCancellation: true,
                            noiseSuppression: true,
                            sampleRate: 44100,
                        }
                    });
                    this.mediaRecorder = new MediaRecorder(stream);
                    this.mediaRecorder.start();
                    this.isAudioRecording = true;
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

                    // 根據不同錯誤顯示不同的錯誤訊息
                    if (error.name === 'NotReadableError') {
                        this.errorMessage = "無法啟動麥克風，可能有其他應用正在使用麥克風。" + error;
                    } else if (error.name === 'NotAllowedError' || error.name === 'SecurityError') {
                        this.errorMessage = "無法獲取麥克風權限，請檢查是否已允許麥克風權限。";
                    } else {
                        this.errorMessage = "麥克風錄音失敗，請檢查您的裝置或瀏覽器設置。";
                    }
                }
            } else {
                this.errorMessage = "你的裝置不支援錄音功能";
            }
        },

        // 停止錄音
        stopAudioRecording() {
            if (this.mediaRecorder && this.isAudioRecording) {
                this.mediaRecorder.stop();
                this.isAudioRecording = false;
            }
        }
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

audio,
video {
    display: block;
    margin-top: 20px;
}

/* 添加錯誤訊息樣式 */
.error-message {
    color: red;
    font-weight: bold;
    margin-bottom: 10px;
    text-align: center;
}
</style>
