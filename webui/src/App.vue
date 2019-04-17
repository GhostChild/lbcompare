<template>
    <div id="app">
        <p>LittleBakas's 实验室 naobu观感测试 </p>
        <p>通过对比选出你认为画面观感更好的版本V1/V2,请不要重复提交</p>
        <div>
            <el-radio v-model="radio" label="1" border>naobu_V1</el-radio>
            <el-radio v-model="radio" label="2" border>naobu_V2</el-radio>
            <el-button type="primary" @click="submit">提交</el-button>
        </div>
        <ul>
            <li v-for="(i,index) in img" class="preview">
                <img class="perviewimg" :src=i.origin width="160" height="90" @click="changeimg(index)">
            </li>
        </ul>
        <el-menu :default-active="activeIndex" mode="horizontal">
            <el-menu-item index="1" @click="ov1(imageindex)">origin / naobu_V1</el-menu-item>
            <el-menu-item index="2" @click="ov2(imageindex)">origin / naobu_V2</el-menu-item>
            <el-menu-item index="3" @click="v1v2(imageindex)">naobu_V1 / naobu_V2</el-menu-item>
        </el-menu>
        <TwentyTwenty
                :before="beforeimg"
                :before-label="beforelabel"
                :after="afterimg"
                :after-label="afterlabel">
        </TwentyTwenty>
    </div>
</template>


<script>
	import TwentyTwenty from 'vue-twentytwenty'
	import axios from 'axios'


	export default {
		name: 'app',
		components: {
			TwentyTwenty
		},
		data() {
			return {
				beforeimg: "https://i.loli.net/2019/04/18/5cb7641da9d52.png",
				beforelabel: "origin",
				afterimg: "https://i.loli.net/2019/04/18/5cb764a7e9951.png",
				afterlabel: "naobu_V1",

				activeIndex: '1',

				radio: 0,

				imageindex: 0,

				img: [{
					origin: "https://i.loli.net/2019/04/18/5cb7641da9d52.png",
					naobu_V1: "https://i.loli.net/2019/04/18/5cb764a7e9951.png",
					naobu_V2: "https://i.loli.net/2019/04/18/5cb766bd206d2.png",
				},
					{
						origin: "https://i.loli.net/2019/04/18/5cb79420b8bd5.png",
						naobu_V1: "https://i.loli.net/2019/04/18/5cb794340f168.png",
						naobu_V2: "https://i.loli.net/2019/04/18/5cb7943d7accb.png",
					}, {
						origin: "https://i.loli.net/2019/04/18/5cb794480ae26.png",
						naobu_V1: "https://i.loli.net/2019/04/18/5cb79450a946e.png",
						naobu_V2: "https://i.loli.net/2019/04/18/5cb79458e5db7.png"

					}, {
						origin: "https://i.loli.net/2019/04/18/5cb792590a8a2.png",
						naobu_V1: "https://i.loli.net/2019/04/18/5cb77b711b4ad.png",
						naobu_V2: "https://i.loli.net/2019/04/18/5cb77b814fb19.png"

					}, {
						origin: "https://i.loli.net/2019/04/18/5cb792f82f594.png",
						naobu_V1: "https://i.loli.net/2019/04/18/5cb79312322a8.png",
						naobu_V2: "https://i.loli.net/2019/04/18/5cb7931fae170.png"
					}]
			}
		},
		methods: {
			ov1(index) {
				this.beforeimg = this.img[index].origin
				this.beforelabel = "origin"
				this.afterimg = this.img[index].naobu_V1
				this.afterlabel = "naobu_V1"
			},
			ov2(index) {
				this.beforeimg = this.img[index].origin
				this.beforelabel = "origin"
				this.afterimg = this.img[index].naobu_V2
				this.afterlabel = "naobu_V2"
			},
			v1v2(index) {
				this.beforeimg = this.img[index].naobu_V1
				this.beforelabel = "naobu_V1"
				this.afterimg = this.img[index].naobu_V2
				this.afterlabel = "naobu_V2"
			},
			changeimg(i) {
				this.imageindex = i
				this.activeIndex = 1
				this.ov1(i)
			},
			submit() {
				if (this.radio == 0) {
					this.$message.error('请选择一个选项再提交')
				} else {
					axios.get('http://littlebakas.moe/submit', {
						params: {
							radio: this.radio
						}
					}).then(res => {
						console.log(res.data)
						this.$message({
							message: '提交成功',
							type: 'success'
						})
					}).catch(res => {
						console.log(res)
					})
				}
			}
		}
	}
</script>

<style>
    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
        margin-top: 10px;
    }

    .twentytwenty-container .twentytwenty-overlay {
        background: none;
    }

    .preview {
        display: inline-block;
    }

    .perviewimg {
        margin-right: 10px;
        margin-left: 10px;
        /*border: 2px solid #409EFF;*/
        border-radius: 4px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);
    }
</style>
