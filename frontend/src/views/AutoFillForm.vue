<template>
    <div class="auto-fill-form">
        <div class="button-container">
            <button class="button" @click="button1Clicked">Button 1</button>
            <button class="button" @click="button2Clicked">Button 2</button>
            <button class="button" @click="button3Clicked">Button 3</button>
        </div>
    </div>
</template>

<script>
import {PDFDocument, rgb} from 'pdf-lib';
import fontkit from '@pdf-lib/fontkit';

export default {
    methods: {
        async button1Clicked() {
            console.log("Button 1 is pressed");

            try {
                // Fetch the PDF from the API
                const response = await fetch('http://localhost:8081/api/autofillform/1');
                const arrayBuffer = await response.arrayBuffer();

                // Load the PDF using PDF-lib
                const pdfDoc = await PDFDocument.load(arrayBuffer);

                // Register fontkit to handle custom fonts
                pdfDoc.registerFontkit(fontkit); // Register fontkit properly

                // Fetch the custom font that supports Chinese characters (NotoSansSC-Regular.ttf)
                const fontUrl = require('@/assets/NotoSansSC-Medium.ttf'); // Ensure path to the font file is correct
                const fontBytes = await fetch(fontUrl).then(res => res.arrayBuffer());

                // Embed the custom font into the PDF
                const customFont = await pdfDoc.embedFont(fontBytes);

                // Get the first page of the document
                const pages = pdfDoc.getPages();
                const firstPage = pages[0];

                // Add Chinese text "你好嗎? 📍" to the PDF
                firstPage.drawText('你好嗎?', {
                    x: 258,
                    y: 690,
                    size: 20,
                    font: customFont,
                    color: rgb(0, 0, 0),
                });
                firstPage.drawText('0918955462', {
                    x: 258,
                    y: 690-38,
                    size: 20,
                    font: customFont,
                    color: rgb(0, 0, 0),
                });

                // Fetch and embed the image (ensure the path is correct)
                const imageUrl = require('@/assets/sampleimg.png'); // Replace with your image path
                const imageBytes = await fetch(imageUrl).then(res => res.arrayBuffer());
                const embeddedImage = await pdfDoc.embedPng(imageBytes); // Use embedJpg for JPG

                // Draw the image on the PDF
                firstPage.drawImage(embeddedImage, {
                    x: 50,
                    y: 500, // Adjust coordinates as needed
                    width: 100, // Adjust size as needed
                    height: 200,
                });
                

                // Save the modified PDF
                const pdfBytes = await pdfDoc.save();

                // Trigger the download of the modified PDF
                const blob = new Blob([pdfBytes], { type: 'application/pdf' });
                const link = document.createElement('a');
                link.href = URL.createObjectURL(blob);
                link.download = 'modified.pdf';
                link.click();
            } catch (error) {
                console.error('Error modifying PDF:', error);
            }
        },
        button2Clicked() {
            console.log("Button 2 is pressed");
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
