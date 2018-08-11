<template>
    <div>
        <input placeholder="First name" v-model="firstName"/>
        <input placeholder="Last name" v-model="lastName"/>
        <input placeholder="Contractor" v-model="contractor"/>
        <input placeholder="Project" v-model="project"/>
        <button @click="onSave">Save</button>
    </div>
</template>

<script>
export default {
    name: 'CreateUserPage',

    data() {
        return {
            firstName: '',
            lastName: '',
            contractor: '',
            project: ''
        }
    },
    methods: {
        onSave() {
            console.log('[LOG] onSave', this.firstName, this.lastName, this.contractor, this.project);
            this.firstName = '';
            this.lastName = '';
            this.contractor = '';
            this.project = '';
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
