<template>
    <div class="activities-list">
        <h1>活動列表</h1>
        <ul>
            <li v-for="activity in activities" :key="activity.id" class="activity-item">
                <div class="activity-header">
                    <h3>{{ activity.title }}</h3>
                    <p class="activity-time">{{ new Date(activity.time).toLocaleString() }}</p>
                </div>
                <p class="activity-description">{{ activity.description }}</p>
                <p class="activity-location">地點：{{ activity.location }}</p>

                <!-- 使用 v-if 確保只有當圖片 URL 存在時才顯示圖片 -->
                <img v-if="activity.image_url" :src="activity.image_url" alt="Activity Image" class="activity-image">

                <!-- Google Maps Iframe for Location -->
                <div v-if="activity.location" class="map-container">
                    <iframe width="100%" height="300" frameborder="0" scrolling="no" marginheight="0" marginwidth="0"
                        :src="'https://maps.google.com/maps?width=100%25&amp;height=300&amp;hl=zh-TW&amp;q=' + encodeURIComponent(activity.location) + '&amp;t=&amp;z=14&amp;ie=UTF8&amp;iwloc=B&amp;output=embed'">
                    </iframe>
                </div>
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

<style scoped>
.activities-list {
    max-width: 900px;
    margin: 40px auto;
    padding: 30px;
    background: #ffffff;
    border-radius: 12px;
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.1);
}

.activities-list h1 {
    text-align: center;
    font-size: 2rem;
    color: #333;
    margin-bottom: 30px;
}

.activities-list ul {
    list-style-type: none;
    padding: 0;
}

.activity-item {
    background: #f9f9f9;
    border-radius: 12px;
    padding: 20px;
    margin-bottom: 30px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.05);
    transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.activity-item:hover {
    transform: translateY(-5px);
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.1);
}

.activity-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.activity-header h3 {
    font-size: 1.5rem;
    margin: 0;
    color: #0056b3;
}

.activity-time {
    font-size: 0.875rem;
    color: #777;
}

.activity-description {
    margin: 15px 0;
    font-size: 1rem;
    color: #555;
}

.activity-location {
    font-weight: bold;
    color: #333;
}

.activity-image {
    width: 100%;
    max-width: 100%;
    height: auto;
    margin-top: 15px;
    border-radius: 8px;
    object-fit: cover;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.map-container {
    margin-top: 15px;
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.05);
}
</style>
