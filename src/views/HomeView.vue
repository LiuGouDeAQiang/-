<template>
  <div>
    <a :href="imageLink">
      <img :src="require('@/assets/封面.png')" alt="Large Image" style="width: 100%;" />
    </a>
  </div>
</template>

<script>
export default {
  data() {
    return {
      imageLink: 'https://lpl.qq.com/es/lpl/2024/',    // 替换为您的链接URL
    };
  },
};
</script>

<style>
/* 根据需要添加样式 */
div {
  text-align: center;
  margin: 5px; /* 为了美观，可根据需要调整间距 */
}

a {
  text-decoration: none; /* 移除链接的下划线 */
}
</style>
