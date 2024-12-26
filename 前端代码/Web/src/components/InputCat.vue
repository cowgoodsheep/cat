<template>
    <el-form ref="ruleFormRef" :model="cats" :rules="rules" label-width="120px" class="demo-cats" :size="formSize"
        status-icon>
        <div label="姓名" prop="cat_avatarFile" class='upload-box'>
            <el-upload class="avatar-uploader" action="/api/uploadcat" :show-file-list="false"
                :before-upload="beforeAvatarUpload" :on-success="handleAvatarSuccess">
                <img v-if="imageUrl" :src="imageUrl" class="avatar" />
                <el-icon v-else class="avatar-uploader-icon">
                    <Plus />
                </el-icon>
            </el-upload>
        </div>
        <div class="u-title">上传头像</div>
        <el-form-item label="姓名" prop="cat_name">
            <el-input v-model="cats.cat_name" autocomplete="on" />
        </el-form-item>
        <el-form-item label="品种" prop="cat_type">
            <el-input v-model="cats.cat_type" />
        </el-form-item>
        <el-form-item label="性别" prop="cat_gender">
            <el-radio-group v-model="cats.cat_gender">
                <el-radio label="母" />
                <el-radio label="公" />
            </el-radio-group>
        </el-form-item>
        <el-form-item label="生日" required>
            <el-col :span="11">
                <el-form-item prop="cat_date">
                    <el-date-picker v-model="cats.cat_date" type="date" label="Pick a date" placeholder="Pick a date"
                        format="YYYY/MM/DD" value-format="YYYY-MM-DD" style="width: 100%" />
                </el-form-item>
            </el-col>
        </el-form-item>
        <el-form-item label="体重(kg)" prop="cat_weight">
            <el-input-number v-model="cats.cat_weight" :min="0" :max="100" />
        </el-form-item>
        <el-form-item label="是否被领养" prop="cat_adapt">
            <el-switch v-model="cats.cat_adapt" />
        </el-form-item>
        <el-form-item label="领养人编号" prop="cat_owner" :min="0">
            <el-input-number v-model="cats.cat_owner" />
        </el-form-item>
        <el-form-item label="是否绝育" prop="cat_sterilization">
            <el-switch v-model="cats.cat_sterilization" />
        </el-form-item>

        <el-form-item>
            <el-button type="primary" @click="submitForm(ruleFormRef)">
                创建
            </el-button>
            <el-button @click="resetForm(ruleFormRef)">重置</el-button>
        </el-form-item>
    </el-form>
</template>
  
<script lang="ts" setup>
import { reactive, ref } from 'vue'
import type { FormInstance, FormRules, UploadProps } from 'element-plus'
import { ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { get, post } from '@/api';
import { AxiosHeaders } from 'axios'

const imageUrl = ref<string>('');
const emit = defineEmits();
interface CatInfo {
    cat_name: string,
    cat_type: string,
    cat_gender: string,
    cat_date: string,
    cat_weight: string,
    cat_adapt: boolean,
    cat_sterilization: boolean,
    cat_owner: number | null,
    cat_avatar: string,
}

const formSize = ref('default')
const ruleFormRef = ref<FormInstance>()

const cats = reactive<CatInfo>({
    cat_name: '',
    cat_type: '',
    cat_gender: '',
    cat_date: '',
    cat_weight: '',
    cat_adapt: false,
    cat_sterilization: false,
    cat_owner: null,
    cat_avatar: '',
})

const rules = reactive<FormRules<CatInfo>>({
    cat_name: [
        { required: true, message: '请输入猫猫的名字', trigger: 'blur' },
        { min: 0, max: 20, message: '最多20字', trigger: 'blur' },
    ],
    cat_type: [
        { required: true, message: '请输入猫猫的品种', trigger: 'blur' },
        { min: 0, max: 20, message: '最多20字', trigger: 'blur' },
    ],
    cat_gender: [
        {
            required: true,
            message: '请选择猫猫的性别',
            trigger: 'change',
        },
    ],
    cat_date: [
        {
            type: 'date',
            required: true,
            message: '请输入猫猫的出生日期',
            trigger: 'change',
        },
    ],

})
const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {
    if (rawFile.type !== 'image/jpeg') {
        ElMessage.error('Avatar picture must be JPG format!')
        return false
    } else if (rawFile.size / 1024 / 1024 > 10) {
        ElMessage.error('Avatar picture size can not exceed 10MB!')
        return false
    }
    return true
}
const handleAvatarSuccess: UploadProps['onSuccess'] = (
    response,
    uploadFile
) => {
    imageUrl.value = URL.createObjectURL(uploadFile.raw!)
    cats.cat_avatar = response.data.url
}

const submitForm = async (formEl: FormInstance | undefined) => {
    console.log(formEl);
    if (!formEl) return
    await formEl.validate(async (valid, fields) => {
        if (valid) {
            console.log('submit!')
            try {
                console.log(cats, 11111)
                const result = await post('/catm', cats);
                console.log(result);
                ElMessage.success("创建成功！")
            } catch (error) {
                console.error('Error fetching data:', error);
                ElMessage.error("创建失败！")
            }
            emit("catCreated")
        } else {
            console.log('error submit!', fields)
        }
    })
    formEl.resetFields()
    imageUrl.value = ''
}

const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
    imageUrl.value = ''
}

</script>

<style scoped>
.upload-box {
    display: flex;
    justify-content: center;
    align-content: space-between;
    margin: 15px;
}

.u-title {
    display: flex;
    justify-content: center;
    margin: 15px;
}

.avatar-uploader .avatar {
    width: 178px;
    height: 178px;
    display: block;
}
</style>

<style>
.avatar-uploader .el-upload {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
    border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    text-align: center;
}
</style>