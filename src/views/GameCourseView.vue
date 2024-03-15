<template>
  <div>
    <h1>赛程详情</h1>
    <div class="calendar">
      <div class="calendar-header">
        <button v-for="date in calendarDates" :key="date" @click="selectDate(date)" :class="{ 'selected-date': selectedDate === date }">{{ date }}</button>
      </div>
    </div>
    <table class="custom-table">
      <thead>
      <tr>
        <th class="table-header">状态</th>
        <th class="table-header">赛事名称</th>
        <th class="table-header">赛事阶段</th>
        <th class="table-header">赛事时间</th>
        <th class="table-header">对阵情况</th>
        <th class="table-header">相关内容</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="item in filteredData" :key="item.ID">
        <td class="table-cell">{{ getStatusText(item.MatchStatus) }}</td>
        <td class="table-cell">{{ item.GameName }}</td>
        <td class="table-cell">{{ item.GameTypeName }}</td>
        <td class="table-cell">{{ item.MatchDate }}</td>
        <td class="table-cell">{{ item.BMatchName }}</td>
        <td class="table-cell">暂无</td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      gameData: [],
      selectedDate: null,
    };
  },
  mounted() {
    const game = this.$route.query.game;
    this.fetchGameData(game);
    // 设置当前日期为默认选中日期
    this.selectedDate = this.getCurrentDate();
  },
  computed: {
    calendarDates() {
      // 从 gameData 中获取所有日期并去重
      return [...new Set(this.gameData.map(item => item.MatchDate.substring(0, 10)))];
    },
    filteredData() {
      // 根据 selectedDate 过滤 gameData
      if (this.selectedDate) {
        return this.gameData.filter(item => item.MatchDate.substring(0, 10) === this.selectedDate);
      } else {
        return this.gameData;
      }
    },
  },
  methods: {
    getStatusText(status) {
      // 状态文本转换
      if (status === "1") {
        return "未开始";
      } else if (status === "2") {
        return "正在进行中";
      } else if (status === "3") {
        return "已结束";
      } else {
        return "";
      }
    },
    fetchGameData(game) {
      axios.get(`/api/gamesCourse?game=${game}`)
          .then(response => {
            this.gameData = response.data.data;
          })
          .catch(error => {
            console.error('发生错误:', error);
          });
    },
    selectDate(date) {
      this.selectedDate = date;
    },
    getCurrentDate() {
      const currentDate = new Date();
      const year = currentDate.getFullYear();
      const month = String(currentDate.getMonth() + 1).padStart(2, '0');
      const day = String(currentDate.getDate()).padStart(2, '0');
      return `${year}-${month}-${day}`;
    },
  }
};
</script>

<style>
.custom-table {
  border-collapse: collapse;
  margin-left: 450px;
  font-size: 20px;
}

.table-header {
  padding: 20px;
}

.table-cell {
  padding: 25px;
}

/* 设置字段之间的间距 */
.table-header + .table-header,
.table-cell + .table-cell {
  margin-left: 50px;
}

/* 设置每一条数据之间的间距 */
.table-cell:first-child {
  margin-top: 20px;
}

.calendar {
  margin-top: 20px;
  display: flex;
  flex-wrap: wrap;
}

.calendar-header {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.selected-date {
  background-color: forestgreen;
  color: #ffffff;
}
</style>