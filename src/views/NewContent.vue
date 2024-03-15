<template>
  <div>
    <h3>{{ article.Title }}</h3>
    <p>{{ article.Content }}</p>

    <h2>讨论区</h2>
    <ul>
      <li v-for="comment in comments" :key="comment.CommentUid">
        <p>{{ comment.Content }}</p>
        <button @click="toggleReply(comment)">回复</button>
        <div v-if="showReply[comment.CommentUid]">
          <input type="text" v-model="replyContent[comment.CommentUid]">
          <button @click="postComment(replyContent[comment.CommentUid], comment.CommentUid)">提交回复</button>
        </div>
        <ul>

          <li v-for="reply in comment.Children" :key="reply.CommentUid">
            <p>➤ {{ reply.Content }}</p> <!-- 用箭头符号表示子评论 -->
            <button @click="toggleReply(reply)">回复</button>
            <div v-if="showReply[reply.CommentUid]">
              <input type="text" v-model="replyContent[reply.CommentUid]">
              <button @click="postComment(replyContent[reply.CommentUid], comment.CommentUid)">提交回复</button>
            </div>
          </li>
        </ul>
      </li>
    </ul>

    <div>
      <h3>发表评论</h3>
      <input type="text" v-model="commentContent">
      <button @click="postComment(commentContent)">提交评论</button>
    </div>
  </div>
</template>
<script>
import axios from "axios";
import { v4 as uuidv4 } from 'uuid';
export default {
  data() {
    return {
      article: {},
      comments: [],
      showReply: {},
      replyContent: {},
      commentContent: ""
    };
  },
  mounted() {
    const titleUid = this.$route.query.titleUid;
    this.getArticleDetails(titleUid);
    this.getComments(titleUid);
  },
  methods: {
    getArticleDetails(titleUid) {
      axios
          .get(`/api/textContent?titleUid=${titleUid}`)
          .then(response => {
            this.article = response.data.data;
            console.log('文章内容',this.article)
          })
          .catch(error => {
            console.error(error);
          });
    },
    generateUniqueCommentUid() {
      return uuidv4();
    },
    getComments(titleUid) {
      axios
          .get(`/api/commentGet?titleUid=${titleUid}`)
          .then(response => {
            this.comments = response.data.data;
            console.log("获取的评论",this.comments)
          })
          .catch(error => {
            console.error(error);
          });
    },

    toggleReply(item) {
      this.$set(this.showReply, item.CommentUid, !this.showReply[item.CommentUid]);
    },
    postComment(content, parentUid = "") {
      const titleUid = this.$route.query.titleUid;
      const userUid = localStorage.getItem("userId");
      const commentUid = this.generateUniqueCommentUid();
      const createdTime = new Date();

      const comment = {
        CommentUid: commentUid,
        Content: content,
        UserUid: userUid,
        ParentUid: parentUid,
        Children: [],
        TitleUid: titleUid,
        CreatedTime: createdTime,
        FloorUid:"0",

      };

      axios
          .post(`/api/commentPost`, comment)
          .then(response => {
            if (parentUid) {
              this.$set(this.replyContent, parentUid, ""); // 清空回复输入框
            } else {
              this.commentContent = ""; // 清空评论输入框
            }
            const newComment = response.data; // 获取新评论或回复
            if (parentUid) {
              const parentComment = this.comments.find(item => item.CommentUid === parentUid);
              parentComment.Children.push(newComment);
            } else {
              this.comments.push(newComment); // 将新评论或回复添加到评论列表
            }
          })
          .catch(error => {
            console.error(error);
          });
    }
  }
};
</script>