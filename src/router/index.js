import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import CreateView from '@/views/CreateView.vue'
import GameView from '@/views/GameView.vue'
import NewView from "@/views/NewView.vue";
import PlayerView from "@/views/PlayerView.vue";
import RefereeView from "@/views/RefereeView.vue";
import PlaybackView from "@/views/PlaybackView.vue";
import TeamView from "@/views/TeamView.vue";
import CourseView from "@/views/GameCourseView.vue";
import TeamPlayerView from "@/views/TeamPlayerView.vue";
import NewContent from "@/views/NewContent.vue";
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView
  },
  {
    path: '/gamesCourse',
    name: 'GamesCourse',
    component: CourseView
  },
  {
    path: '/create',
    name: 'create',
    component: CreateView
  },
  {
    path: '/game',
    name: 'game',
    component: GameView
  },
  {
    path: '/player',
    name: 'player',
    component: PlayerView
  },
  {
    path: '/playback',
    name: 'playback',
    component: PlaybackView
  },
  {
    path: '/referee',
    name: 'referee',
    component: RefereeView
  },
  {
    path: '/team',
    name: 'team',
    component: TeamView
  },{
    path: '/teamPlayer',
    name: 'teamPlayer',
    component: TeamPlayerView
  },
  {
    path: '/New',
    name: 'New',
    component: NewView
  },{
    path: '/NewContent',
    name: 'NewContent',
    component: NewContent
  },

]

const router = new VueRouter({
  routes
})

export default router
