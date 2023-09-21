const fileUploadInput = document.querySelector('#upload_file') as HTMLInputElement
const fileUploadLabelText = (Array as any).from(document.querySelector('#upload_file_label')?.childNodes as any)
    .find(
        (x:any)=>x.nodeName === '#text' && x.textContent.trim()
    )

fileUploadInput?.addEventListener('change', (event: any) => {
    const file = (fileUploadInput.files ?? [])[0]
    fileUploadLabelText.textContent = file.name
    console.log(file)
})

const params = new URLSearchParams(window.location.search)
const file_param = params.get('file')

if (file_param) {
    const label = document.querySelector('#upload_file_label') as HTMLElement | any
    label.innerHTML = `
        <a id="download_file_a" href="/file/${file_param}" target="_blank">${file_param}</a>
    `
    
    const a = document.querySelector('#download_file_a') as HTMLElement | any

    const btn = document.querySelector('#file_upload_button')
    btn?.addEventListener('click', (e)=> {
        a.click()
    });
    (btn as any).innerHTML = 'Download'

    const form = document.querySelector('#file')
    form?.addEventListener('submit', (e)=> {
        e.preventDefault()
    })
}