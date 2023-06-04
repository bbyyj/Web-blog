import Vue from "vue";
import VueRouter from "vue-router";

const Index = () => import("../pages");
const Home = () => import("../pages/Home.vue");
const BlogDetail = () => import("../pages/BlogDetail.vue");
const ChooseQS = () => import("../pages/ChooseQS.vue");
const StudyTest = () => import("../pages/StudyTest.vue");
const HelpChoose = () => import("../pages/HelpChoose.vue");
const AskBoard = () => import("../pages/AskBoard");
const ResourceLib = () => import("../pages/ResourceLib");
const NotFound = () => import("../components/error/NOTFOUND.vue");
const InternalError = () => import("../components/error/INTERNALERROR.vue");
const CourseDetail = () => import("../components/CourseDetail.vue");

const Login = () => import("../pages/admin/Login");
const AdminHome = () => import("../pages/admin/Home");
const Welcome = () => import("../pages/admin/Welcome.vue");
const ListBlogs = () => import("../pages/admin/blogs/ListBlogs.vue");
const AddBlog = () => import("../pages/admin/blogs/AddBlog.vue");
const ManageMessage = () => import("../pages/admin/messages/ManageMessage");
const ListMottos = () => import("../pages/admin/blogs/listMottos");
const ManageResLink = () => import("../pages/admin/resourecLink/manageResLink");
const ManageMusic = () => import("../pages/admin/music/ManageMusic");
const ManageTest = () => import("../pages/admin/StudyTest/ManageTest");
const ManageSandC = () => import("../pages/admin/StudyTest/ManageSandC");
const ManagehaveAsked = () => import("../pages/admin/Ask/ManagehaveAsked.vue");
const ManagenoAsk = () => import("../pages/admin/Ask/ManagenoAsk.vue");
const ManageComments = () =>
  import("../pages/admin/ChooseSubject/ManageComments.vue");
const ManageSubjects = () =>
  import("../pages/admin/ChooseSubject/ManageSubjects.vue");

const ListTags = () => import("../pages/admin/tags/ListTags.vue");
const ListTypes = () => import("../pages/admin/types/ListTypes.vue");

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: Index,
    meta: {
      auth: false,
    },
    redirect: "/home",
    children: [
      { path: "/home", component: Home, meta: { auth: false } },
      { path: "/blogDetail", component: BlogDetail, meta: { auth: false } },
      { path: "/StudyTest", component: StudyTest, meta: { auth: false } },
      { path: "/ChooseQS", component: ChooseQS, meta: { auth: false } },
      { path: "/HelpChoose", component: HelpChoose, meta: { auth: false } },
      { path: "/AskBoard", component: AskBoard, meta: { auth: false } },
      { path: "/resourceLib", component: ResourceLib, meta: { auth: false } },
      { path: "/notFound", component: NotFound, meta: { auth: false } },

      {
        path: "/internalError",
        component: InternalError,
        meta: { auth: false },
      },
      { path: "/CourseDetail", component: CourseDetail, meta: { auth: false } },
    ],
  },
  { path: "/login", component: Login },
  {
    path: "/manageHome",
    component: AdminHome,
    meta: {
      auth: false,
    },
    redirect: "/welcome",
    children: [
      { path: "/welcome", component: Welcome, meta: { auth: true } },
      { path: "/listBlogs", component: ListBlogs, meta: { auth: true } },
      { path: "/addBlog", component: AddBlog, meta: { auth: true } },
      { path: "/listMottos", component: ListMottos, meta: { auth: true } },
      { path: "/manageMsg", component: ManageMessage, meta: { auth: true } },
      { path: "/manageLink", component: ManageResLink, meta: { auth: true } },
      { path: "/manageMusic", component: ManageMusic, meta: { auth: true } },
      { path: "/manageTest", component: ManageTest, meta: { auth: true } },
      {
        path: "/manageSubjectsAndChapters",
        component: ManageSandC,
        meta: { auth: true },
      },
      {
        path: "/managehaveAsked",
        component: ManagehaveAsked,
        meta: { auth: true },
      },
      { path: "/managenoAsk", component: ManagenoAsk, meta: { auth: true } },
      {
        path: "/manageComments",
        component: ManageComments,
        meta: { auth: true },
      },
      {
        path: "/ListTags",
        component: ListTags,
        meta: { auth: true },
      },
      {
        path: "/ListTypes",
        component: ListTypes,
        meta: { auth: true },
      },
      {
        path: "/manageSubjects",
        component: ManageSubjects,
        meta: { auth: true },
      },
    ],
  },
];

const router = new VueRouter({
  routes,
});
/*
//挂载路由导航守卫,此守卫是用来拦截页面访问的，如果没有token，则不能会被重定向到登录页
// 访问博客页面不需要token，直接放行
// 在登陆时不需要token，直接放行
router.beforeEach((to, from, next) => {
  // to: 将要访问的路径
  // form: 从哪个路径跳转而来
  // next 是一个函数，表示放行
  // next() 放行 next("/login") 强制跳转
  // 如果访问的不是管理员或登录页面，直接放行
  if (!to.meta.auth) {
    return next();
  }

  // 访问的是管理员页面，查看有没有token，如果有，则放行，否则，跳转到登录页
  const tokenStr = window.sessionStorage.getItem("token");
  if (!tokenStr) return next("/login");
  next();
});
*/

export default router;
