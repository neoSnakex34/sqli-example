<script>
    export default {
        methods: {
            GoBack() {
                this.$router.push('/');
            }
        }
    }
</script>

<template>
    <div class="d-flex flex-column" style="width: 30%; margin: auto; padding-top: 3%;">
        <h1 class="text-center fw-bold">Error!</h1>
        <!-- todo verbosely log error to hint for attacks -->
        <p class="text-center fw-bold">An error occurred</p>
        <button @click="GoBack" class="btn fw-bold" style="background-color: var(--Nord-Red); color: var(--Nord-White);">Go back to login</button>
    </div>
</template>