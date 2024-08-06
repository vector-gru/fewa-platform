document.addEventListener('DOMContentLoaded', () => {
    const addVideoBtn = document.getElementById('add-video-btn');
    const videoContainer = document.getElementById('video-container');

    addVideoBtn.addEventListener('click', () => {
        const videoGroup = document.createElement('div');
        videoGroup.className = 'form-group video-group';
        videoGroup.innerHTML = `
            <label for="video-upload">Video Upload</label>
            <input type="file" id="video-upload" name="videos[]" required>

            <label for="video-title">Video Title</label>
            <input type="text" id="video-title" name="videoTitle[]" required>

            <label for="video-description">Video Description</label>
            <textarea id="video-description" name="videoDescription[]" rows="2" required></textarea>

            <label for="video-duration">Video Duration (in minutes)</label>
            <input type="number" id="video-duration" name="videoDuration[]" required>

            <label for="thumbnail">Thumbnail Image</label>
            <input type="file" id="thumbnail" name="thumbnail[]" required>
        `;
        videoContainer.appendChild(videoGroup);
    });
});
