<script>
  export default {
    name: 'App',
    components: {
      
    }
  }
</script>

<template>
    <!-- heading -->
    <div class="d-flex justify-content-center" style="background-color: var(--Nord-Black);">
      <h4 style="color: white; font-weight: bold; padding-top: 2px;">Very Secure WebApp</h4>    
    </div>
    <router-view></router-view>
</template>

<style>
  :root {
    --Nord-Black: #2E3440;
    --Nord-DarkGray: #3B4252;
    --Nord-LightGray: #4C566A;
    --Nord-DarkWhite: #D8DEE9;
    --Nord-White: #ECEFF4;
    --Nord-Blue: #5E81AC;
    --Nord-Green: #A3BE8C;
    --Nord-Red: #BF616A;
  }
  body {
    background-color: var(--Nord-DarkWhite)
  }
</style>