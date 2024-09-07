<template>
    <div class="app">
        <h1>測試與 Flutter WebView 的雙向通訊</h1>

        <!-- 發送訊息到 Flutter -->
        <div>
            <button @click="sendMessageToFlutter">發送訊息到 Flutter</button>
        </div>

        <!-- 顯示來自 Flutter 的回應 -->
        <div v-if="flutterResponse">
            <p>來自 Flutter 的回應：</p>
            <pre>{{ flutterResponse }}</pre> <!-- 美化顯示回應 -->
        </div>

        <!-- 顯示所有的 debug 日誌 -->
        <div>
            <h2>Debug 日誌：</h2>
            <ul>
                <li v-for="(log, index) in logs" :key="index">{{ log }}</li>
            </ul>
        </div>
    </div>
</template>

<script>
import { useConnectionMessage } from '../composables/useConnectionMessage';

export default {
    data() {
        return {
            flutterResponse: '', // 儲存來自 Flutter 的回應
            logs: [] // 儲存所有的 debug 日誌
        };
    },
    methods: {
        // 發送訊息到 Flutter
        sendMessageToFlutter() {
            const messageData = { name: 'userinfo', data: null };
            this.addLog(`即將發送訊息到 Flutter: ${JSON.stringify(messageData)}`);

            useConnectionMessage(messageData.name, messageData.data);  // 使用 useConnectionMessage 發送訊息

            this.addLog("訊息已發送至 Flutter");
        },
        // 添加日誌方法
        addLog(log) {
            this.logs.push(log);
        }
    },
    mounted() {
        // 監聽來自 Flutter 的消息
        window.addEventListener("message", (event) => {
            this.addLog(`收到來自 Flutter 的消息: ${event.data}`);

            try {
                // 測試直接顯示資料，檢查格式
                this.addLog(`原始資料: ${event.data}`);

                const responseData = typeof event.data === "string" ? JSON.parse(event.data) : event.data;

                this.flutterResponse = JSON.stringify(responseData, null, 2); // 美化輸出
                this.addLog(`解析後的消息: ${JSON.stringify(responseData)}`);
            } catch (e) {
                this.addLog(`解析來自 Flutter 的消息失敗: ${e}`);
            }
        });
    }
};
</script>

<style scoped>
.app {
    max-width: 600px;
    margin: 20px auto;
    padding: 20px;
    background-color: #f9f9f9;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

h1 {
    text-align: center;
    margin-bottom: 20px;
}

button {
    display: block;
    margin: 10px auto;
    padding: 10px 20px;
    background-color: #0056b3;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

button:hover {
    background-color: #004494;
}

pre {
    white-space: pre-wrap;
    word-wrap: break-word;
    background-color: #f0f0f0;
    padding: 10px;
    border-radius: 4px;
}

ul {
    padding-left: 20px;
}

li {
    margin-bottom: 5px;
    color: #555;
}
</style>
