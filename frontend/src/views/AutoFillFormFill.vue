<template>
    <div>
      <h1>{{ formType }}</h1>
      <p>顯示 {{ formType }} 對應的資料</p>
    </div>
    <!-- a form create ref to formdata-->
    <form @submit.prevent="submitForm">
      <div v-for="(value, key) in formDetail.objects" :key="key">
        <label :for="key">{{ value.data }}</label>
        <!-- a form if type is image than use upload file -->
        <!-- if value.datatype is image and userData[value.data] have text not null value than here place text, else place file upload-->
        <input v-if="value.datatype === 'image' && userData[value.data]" type="text" :id="key" :name="key" :value="userData[value.data]" v-model="userData[value.data]"/>
        <input v-else-if="value.datatype === 'image'" type="file" :id="key" :name="key" @change="onFileChange"/>
        <!-- a form if type is text than use input text -->
        <input v-else type="text" :id="key" :name="key" :value="userData[value.data]" v-model="userData[value.data]"/>
        

      </div>
      <!-- summit button send data to backend-->
        <button type="submit" >Submit</button>
    </form>
    <button @click="printUserData">print User Data</button>
  </template>
  
  <script>
  export default {
    data() {
      return {
        formTitle: '',
        formType: this.$route.params.formType,
        formDetail: {},
        userData: {},
        userimg: null,
        userid: 5,
      };
    },
    created() {
      this.loadFormData();
      this.loadUserData();
    },
    watch: {
      '$route.params.formType': 'loadFormData'
    },
    methods: {
        onFileChange(e) {
            this.userimg = e.target.files[0];
        },
        async loadFormData() {
            this.formType = this.$route.params.formType;
            try {
                const response = await fetch(`http://localhost:8081/api/autofillform/detail/${this.formType}`);
                
                if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
                }
                console.log(response);
                let data = await response.json();
                // turn "data" into json object
                data = JSON.parse(data.detail);

                console.log(data);
                this.formDetail = data;
            } catch (error) {
                console.error('Error fetching data:', error);
                // Handle error appropriately
            }
        },
        // load user data from backend
        async loadUserData() {
            try {
                const response = await fetch('http://localhost:8081/api/users/' + this.userid);
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                console.log(response);
                let data = await response.json();
                // turn "data" into json object
                console.log(data);
                this.userData = data;
            } catch (error) {
                console.error('Error fetching data:', error);
                // Handle error appropriately
            }
        },
        // print user data
        printUserData() {
            console.log(this.userData);
        },
        // turn user input data into json object and use Put method send to backend
        async submitForm() {
            try {
                // upload image
                if (this.userimg) {
                    let formData = new FormData();
                    formData.append('image', this.userimg);
                    const response = await fetch('http://localhost:8081/api/upload/image', {
                        method: 'POST',
                        body: formData,
                    });
                    if (!response.ok) {
                        throw new Error('Network response was not ok');
                    }
                    let data = await response.json();
                    if (data.path) {
                        console.log("Image uploaded, URL:", data.path);
                        // upload/girl.png remove upload/
                        
                        console.log(data.path.slice(8));
                        const newpath = data.path.slice(8);
                        
                        this.userData.imgurl = newpath;
                        console.log(this.userData);
                        
                        // this.userData.imgurl = data.path ;
                    } else {
                        throw new Error('Failed to upload image. No image URL provided.');
                    }
                }
                    
                
                // Option 2: Using Fetch API
                console.log(this.userData);
                const response = await fetch('http://localhost:8081/api/users/' + this.userid, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(this.userData),
                });
                // go to page /autofillform/download with two get parameters formname and userid
                this.$router.push(`/autofillform/download?formname=${this.formType}&userid=${this.userid}`);
                

                
            
            console.log('Form submitted successfully:', response.data);
            } catch (error) {
                console.error('Error submitting form:', error);
            }
            

        }

    }
  }
  </script>

<style scoped>
    form {
        display: flex;
        flex-direction: column;
        gap: 10px;
    }
    label {
        font-weight: bold;
    }
    input {
        padding: 5px;
    }
    button {
        padding: 10px 20px;
        font-size: 16px;
    }


</style>

  