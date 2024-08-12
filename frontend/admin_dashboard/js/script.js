document.querySelectorAll('.collapsible > a').forEach(item => {
    item.addEventListener('click', function() {
        // Close any other open collapsible content
        document.querySelectorAll('.collapsible.active').forEach(activeItem => {
            if (activeItem !== this.parentElement) {
                activeItem.classList.remove('active');
            }
        });

        // Toggle the clicked collapsible content
        this.parentElement.classList.toggle('active');
    });
});
