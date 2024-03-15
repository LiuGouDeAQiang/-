<template>
  <div>
    <h2>注册</h2>
    <form @submit.prevent="register">
      <label for="name">用户名字:</label>
      <input v-model="name" type="text" id="name" required>

      <label for="password">密码:</label>
      <input v-model="password" type="password" id="password" required>

      <label for="confirmPassword">确认密码:</label>
      <input v-model="confirmPassword" type="password" id="confirmPassword" required>

      <button type="submit">注册</button>

      <div v-if="errorMessage" class="error-message">{{ errorMessage }}</div>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      name: '',
      password: '',
      confirmPassword: '',
      errorMessage: '',
    };
  },
  methods: {
    register() {
      // Check if passwords match
      if (this.password !== this.confirmPassword) {
        this.errorMessage = '密码和确认密码不匹配';
        return;
      }

      // Send registration request to the backend
      axios.post('/register', { name: this.name, password: this.password })
          .then(response => {
            if (response.data.code === 999) {
              // Registration successful
              localStorage.setItem('name', this.name);
              this.$router.push('/'); // Redirect to the main page
            } else {
              // Registration failed, display error message
              this.errorMessage = response.data.message;
              // Clear input fields
              this.name = '';
              this.password = '';
              this.confirmPassword = '';
            }
          })
          .catch(error => {
            console.error('注册失败:', error);
            // Handle error situations
          });
    },
  },
};
</script>
