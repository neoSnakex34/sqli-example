<script>
    export default{
        methods: {
            LogOut() {
                this.$router.push('/');
            }
        }
    }
</script>

<template>
    <div class="d-flex flex-column justify-content-center" style="width: 30%; margin: auto; padding-top: 3%;">
        <h1 class="text-center fw-bold">Success!</h1>
        <p class="text-center fw-bold">You have successfully logged in</p>
        <button @click="LogOut" class="btn fw-bold" style="background-color: var(--Nord-Red); color: var(--Nord-White);">Logout</button>
    </div>
</template>

