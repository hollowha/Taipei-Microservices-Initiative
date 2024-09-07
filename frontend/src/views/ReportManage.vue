<template>
    <div class="report-manage">
        <h1>報告列表</h1>
        <div class="report-list">
            
            <li v-for="report in reports" :key="report.id" class="report-item">
                <span class="report-item-header">
                    <h3>{{ report.title }}</h3>
                    <p class="report-type">{{ report.type }}</p>
                    <p class="report-description">{{ report.description }}</p>
                    <p class="report-time">提交時間：{{ new Date(report.time).toLocaleString() }}</p>
                    <p class="report-id">ID: {{ report.id }}</p>
                    <button class="dlelete-btn" @click="deleteReport(report.id)">刪除</button>
                </span>
            </li>
        </div>
    </div>
</template>

<script>
export default {
    name: 'ReportManage',
    data() {
        return {
            reports: []
        };
    },
    created() {
        this.fetchReports();
    },
    methods: {
        fetchReports() {
            fetch('http://localhost:8081/api/report/all')
                .then(response => response.json())
                .then(data => {
                    this.reports = data;
                })
                .catch(error => {
                    console.error('Error fetching reports:', error);
                });
        },
        deleteReport(id) {
            fetch(`http://localhost:8081/api/report/${id}`, {
                method: 'DELETE',
            })
                .then(response => {
                    if (!response.ok) {
                        throw new Error('Network response was not ok');
                    }
                    return response.json();
                })
                .then(data => {
                    console.log('Report deleted:', data);
                    this.fetchReports();
                })
                .catch(error => {
                    console.error('Error deleting report:', error);
                    alert('刪除報告過程中出現錯誤！');
                });
        }
    }
}

</script>

<style>
.report-manage {
    max-width: 900px;
    margin: 40px auto;
    padding: 30px;
    background: #ffffff;
    border-radius: 12px;
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.1);
}

.report-manage h1 {
    text-align: center;
    font-size: 2rem;
    color: #333;
    margin-bottom: 30px;
}

.report-list {
    list-style-type: none; /* 去除列表點 */
    padding: 0;
    margin: 0;
}

.report-item {
    background: #f9f9f9;
    border-radius: 12px;
    padding: 20px;
    margin-bottom: 30px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    transition: transform 0.2s ease box-shadow 0.3s ease;
}

.report-item-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 8px;
}

.report-type {
    font-size: 0.9rem;
    color: #666;
    font-weight: bold;
}

.report-description {
    font-size: 1rem;
    margin-bottom: 10px;
}

.report-time {
    font-size: 0.8rem;
    color: #999;
}

.delete-btn {
    align-self: flex-end;
    padding: 5px 10px;
    background-color: #ff4d4f;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: background-color 0.3s ease;
}

.delete-btn:hover {
    background-color: #d9363e;
}

.report-type,
.report-description,
.report-time {
    margin: 5px 0; /* 確保各部分之間有適當的間距 */
}

.submit-btn {
    display: block;
    width: 100%;
    padding: 15px;
    background-color: #0056b3;
    color: white;
    border: none;
    border-radius: 6px;
    font-size: 1.25rem;
    font-weight: bold;
    cursor: pointer;
    transition: background-color 0.3s ease, transform 0.3s ease;
}

.submit-btn:hover {
    background-color: #004494;
    transform: translateY(-2px);
}

.submit-btn:active {
    background-color: #003366;
    transform: translateY(0);
}
</style>