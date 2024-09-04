<template>
    <div class="auto-fill-form">
        <div class="button-container">
            <button class="button" @click="button1Clicked">simpleform u=5</button>
            <button class="button" @click="button2Clicked">simpleform u=6</button>
            <button class="button" @click="button3Clicked">Button 3</button>
        </div>
    </div>
</template>

<script>
import {PDFDocument, rgb} from 'pdf-lib';
import fontkit from '@pdf-lib/fontkit';

export default {
    methods: {
        async modifyPdf(userid, textCoordinates={x:258,y:690}, formname) {
            try {
                // Fetch the PDF template
                const response = await fetch('http://localhost:8081/api/autofillform/'+formname);
                const arrayBuffer = await response.arrayBuffer();

                // Fetch the user data
                const userDataResponse = await fetch('http://localhost:8081/api/users/' + userid);
                const userData = await userDataResponse.json();
                console.log(userData);

                // Fetch the img
                const imgResponse = await fetch('http://localhost:8081/api/users/image/' + userData.imgurl);
                const imageBytes = await imgResponse.arrayBuffer();
                
                

                // Load the PDF using PDF-lib
                const pdfDoc = await PDFDocument.load(arrayBuffer);

                // Embed the image in the PDF
                const embeddedImage = await pdfDoc.embedPng(imageBytes);

                // Register fontkit to handle custom fonts
                pdfDoc.registerFontkit(fontkit);

                // Fetch the custom font
                const fontUrl = require('@/assets/NotoSansSC-Medium.ttf');
                const fontBytes = await fetch(fontUrl).then(res => res.arrayBuffer());
                const customFont = await pdfDoc.embedFont(fontBytes);


                // Get the first page of the document
                const pages = pdfDoc.getPages();
                const firstPage = pages[0];

                // Add the custom text
                firstPage.drawText(userData.name, {
                    x: textCoordinates.x,
                    y: textCoordinates.y,
                    size: 20,
                    font: customFont,
                    color: rgb(0, 0, 0),
                });
                firstPage.drawText(userData.phone, {
                    x: textCoordinates.x,
                    y: textCoordinates.y-38,
                    size: 20,
                    font: customFont,
                    color: rgb(0, 0, 0),
                });

                // Draw the image
                const imageCoordinates = { x: 350, y: 480, width: 100, height: 120 };
                firstPage.drawImage(embeddedImage, {
                    x: imageCoordinates.x,
                    y: imageCoordinates.y,
                    width: imageCoordinates.width,
                    height: imageCoordinates.height,
                });

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
            const textCoordinates = { x: 258, y: 690 };
            const formname = "simpleform";

            this.modifyPdf(userid, textCoordinates, formname);
        },
        button2Clicked() {
            console.log("Button 2 is pressed");
            const userid = 6;
            const textCoordinates = { x: 258, y: 690 };
            
            // const formname = "modified";
            const formname = "simpleform";

            this.modifyPdf(userid, textCoordinates, formname);
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
