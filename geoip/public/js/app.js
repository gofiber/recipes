$('#search').on('click', () => {
    ip = $('#ip').val()
    $('#result').val('')
    $.getJSON('/geo?ip=' + ip, (res) => {
        $('#result').val(JSON.stringify(res, null, 4))
    })
})