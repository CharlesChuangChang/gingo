<template>
  <el-container>
    <el-header>
      <nav-page style="display: flex;"></nav-page>
    </el-header>
    <el-container>
      <el-aside width="200px">
        <AsidePage></AsidePage>
      </el-aside>
      <!-- <el-divider  direction="vertical"></el-divider> -->
      <el-main>
        <router-view>

        </router-view>
      </el-main>
    </el-container>
    <el-footer>Footer</el-footer>
  </el-container>
</template>

<script>
// import home from './views/HomePage.vue'
import NavPage from './views/NavPage.vue';
import AsidePage from './views/AsidePage.vue';
export default {
  data() {
    return {
    }
  },
  name: 'App',
  components: {
    // home,
    NavPage,
    AsidePage
  },

  

  methods:{
    GoToLogin() {
      console.log("button clicked")
      this.$router.push('/login')
    }
  }
}

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /* margin-top: 60px; */
}
.el-header {
    height: 70px;
    /* background-color: aquamarine; */
    display: flex;
    align-items: center;
    justify-content: flex-end;
    width: 100%;
    background-image: url(./assets/logo.jpg) no-repeat;
  }

.el-aside {
    height: calc(100vh - 120px);

}

.el-main {
    padding: 0;
    height: calc(100vh - 120px);
}
.el-footer {
    height: 50px;
}

</style>
