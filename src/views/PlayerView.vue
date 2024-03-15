<template>
  <div>
    <h2>团队成员列表</h2>
    <table class="centered-table">
      <thead>
      <tr>
        <th class="table-header">签名</th>
        <th class="table-header">真实姓名</th>
        <th class="table-header">比赛名称</th>
        <th class="table-header">战队名称</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="member in teamMembers" :key="member.nick_name">
        <td class="table-cell">{{ member.NickName.String }}</td>
        <td class="table-cell">{{ member.RealName.String }}</td>
        <td class="table-cell">{{ member.GameName.String }}</td>
        <td class="table-cell">{{ member.TeamName.String }}</td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<style>
.centered-table {
  margin: 0 auto; /* 居中显示 */
  border-collapse: collapse; /* 合并边框 */
}

.table-header {
  padding: 20px; /* 调整列表头部的内边距 */
  background-color: #f0f0f0; /* 列表头部背景颜色 */
}

.table-cell {
  padding: 30px; /* 调整单元格的内边距 */
  border: 1px solid #ccc; /* 单元格边框样式 */
  font-size: 14px; /* 字号大小调整 */
}
</style>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      teamMembers: []
    };
  },
  mounted() {

    axios.get(`/api/Player?`)
        .then(response => {
          this.teamMembers = response.data.data;
          console.log(this.teamMembers);
        })
        .catch(error => {
          console.error('发生错误:', error);
        });
  }
};
</script>