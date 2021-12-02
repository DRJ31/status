function getDateTime(datetime) {
    const dt = new Date(datetime * 1000)
    const month = (dt.getMonth() + 1 < 10 ? '0' : '') + (dt.getMonth() + 1).toString()
    const date = (dt.getDate() < 10 ? '0' : '') + dt.getDate().toString()
    const hour = (dt.getHours() < 10 ? '0' : '') + dt.getHours().toString()
    const minute = (dt.getMinutes() < 10 ? '0' : '') + dt.getMinutes().toString()
    const second = (dt.getSeconds() < 10 ? '0' : '') + dt.getSeconds().toString()
    return `${dt.getFullYear()}-${month}-${date} ${hour}:${minute}:${second}`
}

function getHms(duration) {
    const day = Math.floor(duration / (3600 * 24))
    const hour = Math.floor((duration - day * 3600 * 24) / 3600)
    const minute = Math.floor((duration - hour * 3600 - day * 3600 * 24) / 60)
    const second = duration - hour * 3600 - minute * 60 - day * 3600 * 24
    return { day, hour, minute, second }
}

function setReason(log) {
    if (log.type === 1) {
        return `<span style="color: red">${log.reason.detail}(${log.reason.code})</span>`
    }
    else if (log.type === 2) {
        return `<span style="color: #198754">${log.reason.detail}(${log.reason.code})</span>`
    }
    else {
        return log.reason.detail
    }
}

function setLogs(name) {
    const logs = JSON.parse(window.sessionStorage.getItem("logs"))
    let content = ""

    for (let l of logs[name]) {
        const { day, hour, minute, second } = getHms(Math.abs(l.duration))
        content += `<tr>
                            <td>${getDateTime(l.datetime)}</td>
                            <td>${setReason(l)}</td>
                            <td>${day}d ${hour}h ${minute}m ${second}s</td>
                        </tr>`
    }

    $("#tbody").html(content)
    $("#site").html(name)
}

function showAlert(message, type) {
    const content = `<div class="alert alert-${type} alert-dismissible" role="alert">
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-exclamation-triangle-fill flex-shrink-0 me-2" viewBox="0 0 16 16" role="img" aria-label="Warning:">
                            <path d="M8.982 1.566a1.13 1.13 0 0 0-1.96 0L.165 13.233c-.457.778.091 1.767.98 1.767h13.713c.889 0 1.438-.99.98-1.767L8.982 1.566zM8 5c.535 0 .954.462.9.995l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 5.995A.905.905 0 0 1 8 5zm.002 6a1 1 0 1 1 0 2 1 1 0 0 1 0-2z"/>
                          </svg>${message}
                            <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
                        </div>`
    $("#loading").html(content)
}

$(document).ready(function () {
    $.get("/api", (data) => {
        const logs = {}
        $("#year").html(new Date().getFullYear())
        $("#main").css("display", "block")
        $("#loading").css("display", "none")
        $("#up").html(data.up)
        $("#total").html(data.total)
        $("#progress").attr("aria-valuenow", `${data.up / data.total * 100}`).css("width", `${data.up / data.total * 100}%`).html(`${Math.ceil(data.up / data.total * 10000) / 100}%`)
        let content = ""
        for (let m of data.monitors) {
            logs[m.name] = m.logs
            content += `<li class="list-group-item d-flex justify-content-between align-items-start">
                    <div class="ms-2 me-auto">
                        <div class="fw-bold">
                            <a class="navbar-brand" href="${m.url}">${m.name}</a>
                            <i class="fas fa-info-circle" data-bs-toggle="modal" data-bs-target="#logModal" onclick="setLogs('${m.name}')"></i>
                        </div>
                        <span style="color: #198754">${Math.ceil(Number(m.ratio) * 100) / 100}%</span>
                    </div>
                    ${m.status === 2 ? `<span class="badge bg-success">Up</span>` : `<span class="badge bg-danger">Down</span>`}
                </li>`
        }
        $("#list").html(content)
        window.sessionStorage.setItem("logs", JSON.stringify(logs))
    })
        .fail((err) => {
            showAlert(err.responseText, "danger")
        })
});