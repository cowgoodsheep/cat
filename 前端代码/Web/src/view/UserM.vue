<template>
    <div class="search-box">
        <el-input class="input-box" v-model="input" placeholder="请输入搜索关键词" size="large" />
        <el-button type="primary" class="search-button" size="large" round @click="handleSearch">
            <el-icon>
                <Search />
            </el-icon>
        </el-button>
        <el-select class="select-box" v-model="selectValue" placeholder="搜索方式" size="large" @change="selectChange">
            <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
    </div>
    <div class="add-box">
        <el-button type="primary" size="large" circle @click="dialogFormVisible = true">
            <el-icon>
                <Plus />
            </el-icon>
        </el-button>
        <el-dialog class="dialog-box" v-model="dialogFormVisible" title="增加新的用户">
            <input-user @userCreated="handleUserCreated" />
        </el-dialog>
        <el-dialog class="Edit-box" v-model="dialogEditVisible" title="修改此用户">
            <edit-user :editItem="editItem" @userEdited="handleUserEdited" />
        </el-dialog>
    </div>
    <div class="UserTable">
        <el-table class="table-box" :data="tableData" style="width: 100%">
            <el-table-column prop="user_avatar" label="头像" width="120">
                <template v-slot="{ row }">
                    <el-avatar :size="50" :src="row.user_avatar" />
                </template>
            </el-table-column>
            <el-table-column prop="user_id" label="编号" width="100" />
            <el-table-column prop="user_name" label="姓名" width="100" />
            <el-table-column prop="user_password" label="密码" width="100" />
            <el-table-column prop="user_gender" label="性别" width="100" />
            <el-table-column prop="user_phone" label="手机号码" width="100" />
            <el-table-column prop="user_age" label="年龄" />
            <el-table-column>
                <template v-slot="{ row }">
                    <el-button type="default" link @click="openEditDialog(row)">
                        <el-icon>
                            <Edit />
                        </el-icon>
                    </el-button>
                    <el-button type="default" link @click="deleteRow(row)">
                        <el-icon>
                            <Delete />
                        </el-icon>
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>
  
<script lang="ts" setup>
import InputUser from "../components/InputUser.vue"
import EditUser from "../components/EditUser.vue"
import { ref, onMounted } from 'vue'
import { get, post, remove } from '@/api';
import { ElMessage } from 'element-plus'
import { AxiosError } from 'axios';
interface UserInfo {
    user_id: number,
    user_name: string,
    user_password: string,
    user_phone: string,
    user_age: string,
    user_gender: string,
    user_avatar: string,
}
const input = ref('')
const selectValue = ref('')
const dialogFormVisible = ref(false)
const dialogEditVisible = ref(false)
const editItem = ref()
const tableData = ref<UserInfo[]>([]);
/*搜索框的内容接口 */
const handleSearch = async () => {
    try {
        const url = '/userm' + '?' + 'method=' + selectValue.value + '&' + 'content=' + input.value
        const result = await get(url);
        console.log(result, 11111111);
        tableData.value = result;

    } catch (error) {
        console.error('Error fetching data:', error);
    }
}
const selectChange = () => {
    console.log(selectValue.value);
}
const openEditDialog = (row: UserInfo) => {
    // console.log(row);
    editItem.value = row; // 使用对象的拷贝以防止直接修改原始数据
    dialogEditVisible.value = true; // 打开编辑弹窗
};
const updateRow = (async () => {
    const updatedData = await get('/userm');
    tableData.value = updatedData;
})
const handleUserCreated = () => {
    dialogFormVisible.value = false; // 关闭对话框
    updateRow();
};
const handleUserEdited = () => {
    dialogEditVisible.value = false; // 关闭对话框
    updateRow();
}
let options = [
    {
        value: "id",
        label: "编号",
    },
    {
        value: "name",
        label: "姓名",
    },
    {
        value: "gender",
        label: "性别",
    },
    {
        value: "age",
        label: "年龄",
    },
]

onMounted(async () => {
    try {
        const result = await get('/userm');
        console.log(result);
        tableData.value = result;

    } catch (error) {
        console.error('Error fetching data:', error);
    }
});

const deleteRow = async (row: UserInfo) => {
    console.log(row, 1111);
    try {
        // 向后端发送删除请求，传递该行的数据
        const result = await remove('/userm', row);
        console.log(result);
        // 处理删除成功的逻辑
        ElMessage.success("删除成功！");
        // 重新获取表格数据（可选，根据实际情况）
        updateRow()
    } catch (error) {
        // 处理删除失败的逻辑
        console.error('Error deleting data:', error);
        ElMessage.error("删除失败！");
    }
};
</script>
<style scoped>
.UserTable {

    .table-box {
        height: 75vh;
        overflow: hidden;
    }
}

.add-box {
    position: absolute;
    right: 60px;
    bottom: 1.5vh;
    z-index: 11;

    .dialog-box {
        right: 60px;
    }

}

.search-box {
    position: relative;
    margin: 10px;

    .input-box {
        left: 23.5vw;
        width: 40vw;
    }

    .search-button {
        position: absolute;
        right: 13vw;
        top: 0;
    }

    .select-box {
        position: absolute;
        left: 12vw;
        width: 10vw;
    }
}

.add-box {
    margin: 10px
}
</style>  