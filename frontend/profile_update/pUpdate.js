document.addEventListener('DOMContentLoaded', () => {
    // Fetch country phone codes
    fetch('https://restcountries.com/v3.1/all')
        .then(response => response.json())
        .then(data => {
            const countryCodes = document.getElementById('country-codes');
            const sortedData = data.sort((a, b) => a.name.common.localeCompare(b.name.common));
            sortedData.forEach(country => {
                const option = document.createElement('option');
                option.value = country.idd.root + (country.idd.suffixes ? country.idd.suffixes[0] : '');
                option.textContent = `${country.flag} ${country.name.common} (${option.value})`;
                countryCodes.appendChild(option);
            });
        })
        .catch(error => console.error('Error fetching country codes:', error));

    // Fetch time zones with offsets
    // Fetch time zones
    fetch('https://worldtimeapi.org/api/timezone')
        .then(response => response.json())
        .then(data => {
            const timeZoneSelect = document.getElementById('time-zone');
            data.forEach(timezone => {
                const option = document.createElement('option');
                option.value = timezone;
                option.textContent = timezone;
                timeZoneSelect.appendChild(option);
            });
        })
        .catch(error => console.error('Error fetching time zones:', error));
});
