<template>
    <div>
        <h1 style="color:#606266;">领养系统处理</h1>
    </div>
    <div class="CatTable">
        <el-table class="table-box" :data="adaptData" style="width: 100%" round>
            <el-table-column prop="adapt_id" label="编号" width="250" />
            <el-table-column prop="adapt_owner" label="用户编号" width="250" />
            <el-table-column prop="adapt_cat" label="猫猫编号" width="250" />
            <el-table-column prop="adapt_status" label="申请状态" />
            <el-table-column>
                <template v-slot="{ row }">
                    <el-button  type="success" v-if="row.adapt_status === '未处理' && role === 'admin'"
                        @click='request("已同意", row)'>
                        同意
                    </el-button>
                    <el-button type="danger" v-if="row.adapt_status === '未处理' && role === 'admin'"
                        @click='request("已拒绝", row)'>
                        拒绝
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>
  
<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from "element-plus"
import { get, put } from "@/api"
import { AxiosError } from "axios"
const input = ref('')
const role = localStorage.getItem("role")
interface AdaptInfo {
    adapt_id: number | null,
    adapt_owner: number | null,
    adapt_cat: number | null,
    adapt_date: string,
    adapt_status: string,
}

const adaptData = ref<AdaptInfo[]>([])

const adaptRequest = ref<AdaptInfo>({
    adapt_id: null,
    adapt_owner: null,
    adapt_cat: null,
    adapt_date: '',
    adapt_status: '未处理',
})
const updateRow = (async () => {
    const updatedData = await get("/adaptm");
    adaptData.value = updatedData.data
})
const request = async (status: string, row: AdaptInfo) => {
    adaptRequest.value.adapt_id = row.adapt_id
    adaptRequest.value.adapt_owner = row.adapt_owner
    adaptRequest.value.adapt_cat = row.adapt_cat
    adaptRequest.value.adapt_date = row.adapt_date
    adaptRequest.value.adapt_status = status
    console.log(adaptRequest.value, "requestadapt");
    try {
        const result = await put("/adaptm", adaptRequest.value);
        console.log(result, "tetsttttt");
        ElMessage.success(`${result.data.msg}`)
        updateRow()
    } catch (error) {
        // 使用类型断言将 error 断言为 AxiosError 类型
        const axiosError = error as AxiosError<{ data: { msg: string } }>;
        console.log(axiosError);
        // 如果是 AxiosError，并且包含响应信息
        if (axiosError.isAxiosError && axiosError.response) {
            const errorMsg = axiosError.response.data.data.msg;
            ElMessage.error(`操作失败：${errorMsg}`);
            // 显示具体的后端错误信息
        } else {
            // 显示通用的错误信息
            ElMessage.error('操作失败：未知错误！');
        }
    }
}
onMounted(async () => {
    try {
        const result = await get("/adaptm");
        console.log(result);
        adaptData.value = result.data
    } catch (error) {
        console.error('Error fetching data:', error);
    }
})

</script>
<style scoped>
.tot-box {
    display: flex;
    justify-content: space-around;
    padding: 30px;
}

.CatTable {
    .table-box {
        height: 75vh;
        overflow: hidden;
    }
}

.dialog-footer button:first-child {
    margin-right: 10px;
}
</style>  