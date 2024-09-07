<template>
    <div class="auto-fill-form">
        
        <div class="button-container">
            <button class="button" @click="button1Clicked">simpleform u=5</button>
            <button class="button" @click="button2Clicked">anotherform u=6</button>
            <button class="button" @click="button3Clicked">Button 3</button>
        </div>
    </div>
</template>

<script>
import {PDFDocument, rgb} from 'pdf-lib';
import fontkit from '@pdf-lib/fontkit';

export default {
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
                link.download = formname + '.pdf';
                link.click();
            } catch (error) {
                console.error('Error modifying PDF:', error);
            }
        },

        button1Clicked() {
            const userid = "5";
            
            const formname = "simpleform";

            this.modifyPdf(userid, formname);
        },
        button2Clicked() {
            console.log("Button 2 is pressed");
            const userid = 6;
            
            
            // const formname = "modified";
            const formname = "anotherform";

            this.modifyPdf(userid, formname);
        },
        button3Clicked() {
            console.log("Button 3 is pressed");
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
