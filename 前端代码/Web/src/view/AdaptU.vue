<template>
    <div class="process-box">
        <h1 style="color:#606266;">我的申请</h1>
        <div class="timeline-box">
            <el-timeline>
                <el-timeline-item v-for="item in   adaptData  " :key="item.adapt_owner" :timestamp="item.adapt_date"
                    placement="top">
                    <el-card v-if="item.adapt_owner == nowid">
                        <h4>你申请收养了猫猫，编号为{{ item.adapt_cat }}</h4>
                        <p>申请状态为:{{ item.adapt_status }}</p>
                    </el-card>
                    <el-card v-else>
                        <h4>你还没有发起申请领养哦</h4>
                    </el-card>
                </el-timeline-item>
            </el-timeline>
        </div>
    </div>
    <div class="apply-box">
        <h1 style="color:#606266;">提交新的申请</h1>
        <div class=request-box>
            <el-form :inline="true" :model="formInline" class="demo-form-inline">
                <el-form-item label="申请猫猫的编号">
                    <el-input-number v-model="formInline.cat" size="large" clearable />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="onSubmit">提交</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { reactive, onMounted, ref } from 'vue'
import { get, post } from "@/api"
import { ElMessage } from "element-plus"
import { AxiosError } from "axios"
const nowid = localStorage.getItem("id")
const formInline = reactive({
    cat: null,
    reason: '',
})
interface AdaptInfo {
    adapt_owner: number | null,
    adapt_cat: number | null,
    adapt_date: string,
    adapt_status: string,
}
const adaptData = ref<AdaptInfo[]>([])
const adaptRequest = ref<AdaptInfo>({
    adapt_owner: null,
    adapt_cat: null,
    adapt_date: '',
    adapt_status: '未处理',
})

const onSubmit = async () => {
    console.log('submit!')
    const idFromLocalStorage: string | null = localStorage.getItem("id");
    adaptRequest.value.adapt_owner = idFromLocalStorage ? parseInt(idFromLocalStorage, 10) : 0;
    const currentDateString = new Date().toLocaleDateString();
    adaptRequest.value.adapt_date = currentDateString
    adaptRequest.value.adapt_cat = formInline.cat
    console.log(adaptRequest.value, 111111111)
    try {
        const result = await post("/adaptu", adaptRequest.value)
        console.log(result)
        ElMessage.success("提交申请成功！")
    } catch (error) {
        // 使用类型断言将 error 断言为 AxiosError 类型
        const axiosError = error as AxiosError<{ data: { msg: string } }>;
        console.log(axiosError);
        // 如果是 AxiosError，并且包含响应信息
        if (axiosError.isAxiosError && axiosError.response) {
            const errorMsg = axiosError.response.data.data.msg;
            ElMessage.error(`提交申请失败：${errorMsg}`);
            // 显示具体的后端错误信息
        } else {
            // 显示通用的错误信息
            ElMessage.error('提交申请失败！');
        }
    }
}

onMounted(async () => {
    try {
        const id = localStorage.getItem('id');
        const url = "/adaptu" + "?id=" + id?.toString()
        const result = await get(url);
        console.log(result);
        adaptData.value = result.data
        console.log(adaptData.value)
    } catch (error) {
        console.error('Error fetching data:', error);
    }
});
</script>

<style>
.apply-box {
    padding: 20px;

    .demo-form-inline .el-input {
        --el-input-width: 220px;
    }
}
</style>