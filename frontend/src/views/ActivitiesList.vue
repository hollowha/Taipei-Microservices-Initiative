<template>
    <div class="activities-list">
        <h1>活動列表</h1>
        <ul>
            <li v-for="activity in activities" :key="activity.id">
                <h3>{{ activity.title }}</h3>
                <p>{{ activity.description }}</p>
                <p>地點：{{ activity.location }}</p>
                <p>時間：{{ activity.time }}</p>
                <img :src="activity.image_url" alt="Activity Image" style="width: 100px; height: auto;">
                <!-- <div style="width: 100%">
                    <iframe width="100%" height="600" frameborder="0" scrolling="no" marginheight="0" marginwidth="0"
                        src="https://maps.google.com/maps?width=100%25&amp;height=600&amp;hl=zh-TW&amp;q=1%20Grafton%20Street,%20Dublin,%20Ireland+(My%20Business%20Name)&amp;t=&amp;z=14&amp;ie=UTF8&amp;iwloc=B&amp;output=embed">
                        <a href="https://www.gps.ie/">gps tracker sport</a>
                    </iframe>
                </div> -->
                <div style="width: 100%">
                    <iframe width="100%" height="600" frameborder="0" scrolling="no" marginheight="0" marginwidth="0"
                        src="https://maps.google.com/maps?width=100%25&amp;height=600&amp;hl=zh-TW&amp;q=桃園市 縣府路 256巷+(My%20Business%20Name)&amp;t=&amp;z=14&amp;ie=UTF8&amp;iwloc=B&amp;output=embed">
                        <a href="https://www.gps.ie/">gps tracker sport</a>
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
            fetch('http://localhost:8081/api/activitys/')
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
    margin: 0 auto;
    padding: 20px;
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
</style>