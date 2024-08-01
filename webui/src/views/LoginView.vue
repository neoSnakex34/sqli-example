
<script>
    export default {
        data: function() {
            return {
                // not used now 
                username: '',
                password: ''

                
            };
        },

        methods: {
            async LogIn() {
                
                try { 
                    let response = await this.$http.post(`${this.$apiurl}/login`, {
                        username: this.username,
                        password: this.password
                    }); 

                    console.log(response.data);
                    // if logged in successfully 
                    // mount another view 
                    this.$router.push('/logged');
                } catch (e) {
                    if (e.response.data){
                        console.log(e.response.data);
                    } else {
                        console.log(e);
                    }
                    this.$router.push('/error');
                }
            }
        }
    }
</script>

<template>
    <div class="d-flex flex-column justify-content-center" style="width: 30%; margin: auto; padding-top: 5%;">
        <input v-model="username" class="form-control rounded-bottom-0 input-box" placeholder="username"></input>
        <input v-model="password" type="password" class="form-control rounded-top-0 input-box" placeholder="password"></input>
        <button @click=LogIn() class="btn mt-2 fw-bold" style="background-color: var(--Nord-Green); color: var(--Nord-White)">Login</button>

    </div>
</template>

<style>
    .input-box {
        background-color: var(--Nord-White);

    }

    .input-box:focus {
        background-color: var(--Nord-White);
        box-shadow: 0 0 0 0.25rem var(--Nord-Green);
        z-index: 1;
    }


</style>