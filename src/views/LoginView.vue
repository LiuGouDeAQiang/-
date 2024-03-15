<template>
  <div>
    <h2>登录</h2>
    <form @submit.prevent="login">
      <label for="uid">用户ID:</label>
      <input v-model="uid" type="text" id="uid" required>

      <label for="password">密码:</label>
      <input v-model="password" type="password" id="password" required>

      <button type="submit">登录</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  data() {
    return {
      uid: '',
      password: '',
      errorMessage: '', // 新增错误信息
    };
  },
  methods: {
    login() {
      // 发送登录请求到后端
      axios.post('/login', { uid: this.uid, password: this.password })
          .then(response => {
            if (response.data.code === 999) {
              // 登录成功
              localStorage.setItem('uid',this.uid)
              this.$router.push('/'); // 跳转到主页面
            } else {
              // 登录失败，显示错误信息
              this.errorMessage = response.data.message;
              // 清空uid和password
              this.uid = '';
              this.password = '';
            }
          })
          .catch(error => {
            console.error('登录失败:', error);
            // 处理错误情况
          });
    },
  },
};
</script>
