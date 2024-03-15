<template>
  <div>
    <h1>新闻文章列表</h1>
    <ul>
      <li v-for="article in articles" :key="article.id">
        <h3 @click="getArticleDetails(article.Uid)">{{ article.Title }}</h3>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      articles: [],         // 存储文章列表
    };
  },
  mounted() {
    this.getArticles();    // 在组件挂载后获取文章列表
  },
  methods: {
    getArticles() {
      // 使用Ajax发送GET请求获取文章列表
      axios
          .get('/api/textList')
          .then(response => {
            this.articles = response.data.data;
            console.log(this.articles)
          })
          .catch(error => {
            console.error(error);  // 捕获并打印错误信息
          });
    },
    getArticleDetails(uid) {
      // 使用Ajax发送GET请求获取文章详情
      console.log(uid)
      localStorage.setItem("userId","654")
      this.$router.push({
        path: '/NewContent',
        query: {
          titleUid: uid
        }
      });
    },
  }
};
</script>