<template>
    <div class="activities-list">
        <h1>活動列表</h1>
        <ul>
            <li v-for="activity in activities" :key="activity.id">
                <h3>{{ activity.title }}</h3>
                <p>{{ activity.description }}</p>
                <p>地點：{{ activity.location }}</p>
                <p>時間：{{ activity.time }}</p>
                <!-- 使用 v-if 確保只有當圖片 URL 存在時才顯示圖片 -->
                <img v-if="activity.image_url" :src="activity.image_url" alt="Activity Image" class="activity-image">
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    name: 'ActivitiesList',
    data() {
        return {
            activities: []
        };
    },
    created() {
        this.fetchActivities();
    },
    methods: {
        fetchActivities() {
            fetch('http://localhost:8081/api/activitys/all')
                .then(response => response.json())
                .then(data => {
                    this.activities = data;
                })
                .catch(error => {
                    console.error('Error fetching activities:', error);
                });
        }
    }
}
</script>

<style>
.activities-list {
    max-width: 800px;
    margin: 20px auto;
    padding: 20px;
    background: #f4f4f4;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.activities-list ul {
    list-style-type: none;
    padding: 0;
}

.activities-list li {
    border-bottom: 1px solid #ccc;
    margin-bottom: 10px;
    padding-bottom: 10px;
}

.activity-image {
    width: 100%;
    max-width: 300px;
    /* 限制圖片大小 */
    height: auto;
    margin-top: 10px;
}
</style>
