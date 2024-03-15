<template>
  <div>
    <h1>战队列表</h1>
    <ul>
      <li v-for="teamName in teamsName" :key="teamName" @click="goToPlay(teamName)" class="team-item">
        {{ teamName }}
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      teamsName: []
    };
  },
  mounted() {
    this.getData();
  },
  methods: {
    getData() {
      axios.get('/api/teamsList')
          .then(response => {
            this.teams = response.data.teams;
            this.teams.forEach(team => {
              this.teamsName.push(team.TeamName);
            });
          })
          .catch(error => {
            console.error('发生错误:', error);
          });
    },
    goToPlay(team) {
      console.log(team)
      this.$router.push({
        path: '/teamPlayer',
        query: {
          team: team
        }
      });
    }
  }
};
</script>
<style>
.team-item {
  font-family: Arial, sans-serif; /* 设置字体 */
  font-size: 20px;
  margin-bottom: 20px; /* 设置下方间距 */
}
</style>