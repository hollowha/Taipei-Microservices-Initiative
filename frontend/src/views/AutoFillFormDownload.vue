<template>
    <div class="auto-fill-form">
        <h1>Auto Fill Form Download</h1>
        <div class="button-container">
            <button class="button" @click="downloadPdf">Download PDF</button>
        </div>
    </div>
</template>

<script>
import {PDFDocument, rgb} from 'pdf-lib';
import fontkit from '@pdf-lib/fontkit';

export default {
    data() {
        return {
            formname: this.$route.query.formname,
            userid: this.$route.query.userid,
        };
    },
    methods: {
        async modifyPdf(userid, formname) {
            try {
                // Fetch the PDF template
                const response = await fetch('http://localhost:8081/api/autofillform/'+formname);
                const arrayBuffer = await response.arrayBuffer();

                // Fetch the user data
                const userDataResponse = await fetch('http://localhost:8081/api/users/' + userid);
                const userData = await userDataResponse.json();
                console.log(userData);
                let a = "name";
                console.log(userData[a]);

                
                const pdfDetailsOringin = await fetch('http://localhost:8081/api/autofillform/detail/' + formname);
                const pdfDetailsfull = await pdfDetailsOringin.json();
                const pdfDetailsText = pdfDetailsfull.detail;
                const pdfDetails = JSON.parse(pdfDetailsText);
                console.log(pdfDetails);

                // Load the PDF using PDF-lib
                const pdfDoc = await PDFDocument.load(arrayBuffer);
                
                // Register fontkit to handle custom fonts
                pdfDoc.registerFontkit(fontkit);

                // Fetch the custom font
                // get NotoSansSC-Medium from google
                const fontUrl = require('@/assets/NotoSansSC-Medium.ttf');
                const fontBytes = await fetch(fontUrl).then(res => res.arrayBuffer());
                const customFont = await pdfDoc.embedFont(fontBytes);


                // Get the first page of the document
                const pages = pdfDoc.getPages();
                const firstPage = pages[0];

                console.log(pdfDetails.objects.length);

                for (let i = 0; i < pdfDetails.objects.length; i++) {
                    const detail = pdfDetails.objects[i];
                    console.log(detail);

                    if (detail.datatype == "text") {
                        console.log("in text area");
                        firstPage.drawText(userData[detail.data], {
                            x: detail.x,
                            y: detail.y,
                            size: detail.size,
                            font: customFont,
                            color: rgb(0, 0, 0),
                        });
                    } else if (detail.datatype == "image") {
                        const imgResponse = await fetch('http://localhost:8081/api/users/image/' + userData[detail.data]);
                        const imageBytes = await imgResponse.arrayBuffer();
                        const embeddedImage = await pdfDoc.embedPng(imageBytes);
                        firstPage.drawImage(embeddedImage, {
                            x: detail.x,
                            y: detail.y,
                            width: detail.width,
                            height: detail.height,
                        });
                    }
                    
                }

                // Save the modified PDF
                const pdfBytes = await pdfDoc.save();
                // Trigger the download of the modified PDF
                const blob = new Blob([pdfBytes], { type: 'application/pdf' });
                const link = document.createElement('a');
                link.href = URL.createObjectURL(blob);
                link.download = formname + '_' + userData.name + '.pdf';
                link.click();

                // sent the modified pdf to backend
                await fetch('http://localhost:8081/api/autofillform/savePDF/' + formname + '_' + userData.name, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/pdf',
                    },
                    body: pdfBytes,
                });
            } catch (error) {
                console.error('Error modifying PDF:', error);
            }
        },

        async downloadPdf() {
            
            const userid = this.userid;
            const formname = this.formname;
            
            console.log(userid);
            console.log(formname);

            this.modifyPdf(userid, formname);
        },
    },
};
</script>

<style scoped>
.auto-fill-form {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh; /* Adjust as needed */
}

.button-container {
    display: flex;
    flex-direction: column;
    gap: 10px; /* Adjust spacing between buttons as needed */
}

.button {
    padding: 10px 20px;
    font-size: 16px;
}
</style>
