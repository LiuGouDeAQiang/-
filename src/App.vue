<template>
  <div id="app">
    <nav>
      <div class="nav-section one">
        <router-link to="/">首页</router-link>
        <router-link to="/game">赛事</router-link>
        <router-link to="/team">战队</router-link>
        <router-link to="/player">职业选手</router-link>
        <router-link to="/referee">裁判</router-link>
        <router-link to="/new">新闻</router-link>
        <router-link to="/playback">回放</router-link>
        <div class="spacer"></div>

        <!-- Updated conditional rendering -->
        <template v-if="!loggedIn">
          <router-link to="/login" class="nav-link">
            <button class="action-button">登录</button>
          </router-link>
          <router-link to="/create" class="nav-link">
            <button class="action-button">注册</button>
          </router-link>
        </template>

        <template v-else>
          <div class="welcome-text">欢迎：{{ username }}</div>
          <button class="action-button" @click.stop="logout">注销</button>
        </template>
      </div>
    </nav>
    <router-view/>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loggedIn: false,
      username: '',
    };
  },
  mounted() {
    // 检查本地存储是否有 uid，如果有则设置登录状态
    const storedUid = localStorage.getItem('uid');
    if (storedUid) {
      this.loggedIn = true;
      this.username = storedUid;
    }
  },
  methods: {
    logout() {
      // 注销，清除本地存储的 uid 并重置登录状态
      localStorage.removeItem('uid');
      this.loggedIn = false;
      this.username = '';
      // 跳转到登录页面
      this.$router.push('/login');
    },
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

.nav-section {
  display: flex;
  align-items: center;
  min-width: 600px; /* 设置最小宽度，避免在屏幕放大时换行 */
}
.spacer {
  flex-grow: 1;
  max-width: 700px; /* Set the maximum width for spacing */
}
.one {

  margin-left: 30px;
  font-size: 1.4em;
}

nav a {
  display: inline-block;
  font-weight: bold;
  color: #2c3e50;
  margin-right: 100px;
  text-decoration: none; /* 移除下划线 */
}

nav a:last-child {
  margin-right: 0;
}

nav a.router-link-exact-active {
  color: #b8872b;
}


.button-container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
}

.action-button {
  padding: 10px 20px;
  font-size: 20px;
  cursor: pointer;
  background-color: transparent; /* 设置为透明背景色 */
  color: #b8872b; /* 暗金色 */
  border: none;

  border-radius: 5px;
}

.spacer {
  width: 50px;
}
</style>