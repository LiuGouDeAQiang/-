<template>
  <div>
    <h1>赛事列表</h1>
    <ul>
      <li v-for="(game, index) in games" :key="index" @click="goToCourse(game)">
      <div>{{game}}</div>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      games: []
    };
  },
  mounted() {
    this.getData();
  },
  methods: {
    getData() {
      axios.get('/api/gameList')
          .then(response => {
            this.games = response.data.data;
            console.log(this.games)
          })
          .catch(error => {
            console.error('发生错误:', error);
          });
    },
    goToCourse(game) {
      console.log(game)
      this.$router.push({
        path: '/gamesCourse',
        query: {
          game: game
        }
      });
    }
  }
};
</script>