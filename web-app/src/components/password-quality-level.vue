<template>
    <div class="password-quality-level__block" :class="securityLevel">
        <div class="password-quality-level__bar"></div>
        <div class="password-quality-level__bar"></div>
        <div class="password-quality-level__bar"></div>
        <div class="password-quality-level__bar"></div>
    </div>
</template>

<script>
export default {
    name: 'password-quality-level',
    props: ['password'],
    methods: {
       
    },
    computed: {
        securityLevel() {
            const level = {}
            level[`level-${this.passwordValue}`] = true;
            return level;
        },
        passwordValue() {
            return this.hasSymbols + this.hasNumbers + this.hasNiceLength;
        },
        hasSymbols() {
            const regex = new RegExp('.*[\\!\\@\\#\\$\\%\\^\\&\\*\\.\\_\\-\\=\\+].*', 'g');
            return this.password.match(regex) ? 1 : 0;
        },
        hasNumbers() {
            const regex = new RegExp('.*[0-9].*', 'g');
            return this.password.match(regex) ? 1 : 0;
        },
        hasNiceLength() {
            return this.password.length > 8 ? 2 : 1;
        },
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.level-1 .password-quality-level__bar:nth-child(1) {
    background: #982f2c;
}
.level-2 .password-quality-level__bar:nth-child(2) {
    background: #fcc529;;
}
.level-3 .password-quality-level__bar:nth-child(3) {
    background: #a7de7a;
}
.level-4 .password-quality-level__bar:nth-child(4) {
    background: #7accde;
}
.password-quality-level__block {
    // width: 100%;
    height: 1em;
    // background: purple
}
.password-quality-level__bar {
    height: .4em;
    margin-bottom: .6em;
    width: 20%;
    display: inline-block;
    border-radius: 10px;
    background: #bcbbb9
}
</style>
