package main

import (
	"fmt"
	"strings"
)

func main() {
	// email body template with a fixed variable OTP value
	_ = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>OTP from Bohubrihi Team</title>
</head>
<body>
    <h1>OTP: %s</h1>
	<span style="font-size:18px;">%s</span>
    <p>This is your one-time password (OTP) for logging in to your account.</p>
</body>
</html>
`
	// OTP value to be replaced in the template
	otp := "123456"

	// replace the %s placeholder in the template with the OTP value
	// body := fmt.Sprintf(template, otp, otp)

	template := `<!DOCTYPE html>
<html xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" lang="en">
  <head>
    <title></title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <style>
      * {
        box-sizing: border-box
      }

      body {
        margin: 0;
        padding: 0
      }

      a[x-apple-data-detectors] {
        color: inherit !important;
        text-decoration: inherit !important
      }

      #MessageViewBody a {
        color: inherit;
        text-decoration: none
      }

      p {
        line-height: inherit
      }

      .desktop_hide,
      .desktop_hide table {
        mso-hide: all;
        display: none;
        max-height: 0;
        overflow: hidden
      }

      .image_block img+div {
        display: none
      }

      @media (max-width:620px) {
        .row-content {
          width: 100% !important
        }

        .mobile_hide {
          display: none
        }

        .stack .column {
          width: 100%;
          display: block
        }

        .mobile_hide {
          min-height: 0;
          max-height: 0;
          max-width: 0;
          overflow: hidden;
          font-size: 0
        }

        .desktop_hide,
        .desktop_hide table {
          display: table !important;
          max-height: none !important
        }
      }
    </style>
  </head>
  <body style="margin:0;background-color:#1e1e1e;padding:0;-webkit-text-size-adjust:none;text-size-adjust:none">
    <table class="nl-container" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;background-color:#1e1e1e">
      <tbody>
        <tr>
          <td>
            <table class="row row-1" align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;background-color:#f6f6f6">
              <tbody>
                <tr>
                  <td>
                    <table class="row-content stack" align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;background-color:#1e1e1e;color:#000;width:600px" width="600">
                      <tbody>
                        <tr>
                          <td class="column column-1" width="100%" style="mso-table-lspace:0;mso-table-rspace:0;font-weight:400;text-align:left;vertical-align:top;border-top:0;border-right:0;border-bottom:0;border-left:0">
                            <table class="image_block block-1" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0">
                              <tr>
                                <td class="pad" style="padding-bottom:20px;padding-left:10px;padding-right:10px;padding-top:20px;width:100%">
                                  <div class="alignment" align="center" style="line-height:10px">
                                    <img src="https://bohubrihi.com/wp-content/uploads/2022/05/Bohubrihi-Online-Courses-2.png" style="display:block;height:auto;border:0;width:146px;max-width:100%" width="146" alt="Bohubrihi" title="Bohubrihi">
                                  </div>
                                </td>
                              </tr>
                            </table>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </td>
                </tr>
              </tbody>
            </table>
            <table class="row row-2" align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;background-color:#f6f6f6">
              <tbody>
                <tr>
                  <td>
                    <table class="row-content stack" align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;background-color:#fff;color:#000;width:600px" width="600">
                      <tbody>
                        <tr>
                          <td class="column column-1" width="100%" style="mso-table-lspace:0;mso-table-rspace:0;font-weight:400;text-align:left;border-left:1px solid #f0f0f0;border-right:1px solid #f0f0f0;padding-bottom:15px;padding-left:15px;padding-right:15px;padding-top:15px;vertical-align:top;border-top:0;border-bottom:0">
                            <table class="text_block block-1" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;word-break:break-word">
                              <tr>
                                <td class="pad">
                                  <div style="font-family:sans-serif">
                                    <div class style="font-size:12px;font-family:Tahoma,Verdana,Segoe,sans-serif;mso-line-height-alt:18px;color:#171717;line-height:1.5">
                                      <p style="margin:0;font-size:12px;mso-line-height-alt:45px">
                                        <span style="font-size:30px;">
                                          <strong>
                                            <span style="font-size:26px;">লগইন করতে কোডটি ইউজ করুন</span>
                                            <br>
                                          </strong>
                                        </span>
                                      </p>
                                    </div>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <table class="text_block block-2" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;word-break:break-word">
                              <tr>
                                <td class="pad">
                                  <div style="font-family:sans-serif">
                                    <div class style="font-size:12px;font-family:Tahoma,Verdana,Segoe,sans-serif;mso-line-height-alt:18px;color:#595959;line-height:1.5">
                                      <p style="margin:0;mso-line-height-alt:27px">
                                        <span style="font-size:18px;">বহুব্রীহি'র নতুন সাইটে আপনার অ্যাকাউন্টটি আপডেট করার জন্য আপনাকে নিচের ভেরিফিকেশন কোডটি ইউজ করতে হবে।</span>
                                      </p>
                                    </div>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <table class="divider_block block-3" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0">
                              <tr>
                                <td class="pad">
                                  <div class="alignment" align="center">
                                    <table border="0" cellpadding="0" cellspacing="0" role="presentation" width="100%" style="mso-table-lspace:0;mso-table-rspace:0">
                                      <tr>
                                        <td class="divider_inner" style="font-size:1px;line-height:1px;border-top:1px solid #bbb">
                                          <span>&#8202;</span>
                                        </td>
                                      </tr>
                                    </table>
                                  </div>
                                </td>
                              </tr>
                            </table>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </td>
                </tr>
              </tbody>
            </table>
            <table class="row row-3" align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;background-color:#f6f6f6">
              <tbody>
                <tr>
                  <td>
                    <table class="row-content stack" align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;background-color:#ddd;color:#000;width:600px" width="600">
                      <tbody>
                        <tr>
                          <td class="column column-1" width="100%" style="mso-table-lspace:0;mso-table-rspace:0;font-weight:400;text-align:left;border-left:1px solid #f0f0f0;border-right:1px solid #f0f0f0;padding-bottom:15px;padding-left:15px;padding-right:15px;padding-top:15px;vertical-align:top;border-top:0;border-bottom:0">
                            <table class="text_block block-1" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;word-break:break-word">
                              <tr>
                                <td class="pad">
                                  <div style="font-family:sans-serif">
                                    <div class style="font-size:12px;font-family:Tahoma,Verdana,Segoe,sans-serif;mso-line-height-alt:18px;color:#595959;line-height:1.5">
                                      <p style="margin:0;text-align:center;mso-line-height-alt:27px">
                                        <span style="font-size:18px;">OTP: %s</span>
                                      </p>
                                    </div>
                                  </div>
                                </td>
                              </tr>
                            </table>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </td>
                </tr>
              </tbody>
            </table>
            <table class="row row-4" align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;background-color:#f6f6f6">
              <tbody>
                <tr>
                  <td>
                    <table class="row-content stack" align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;background-color:#fff;color:#000;width:600px" width="600">
                      <tbody>
                        <tr>
                          <td class="column column-1" width="100%" style="mso-table-lspace:0;mso-table-rspace:0;font-weight:400;text-align:left;border-left:1px solid #f0f0f0;border-right:1px solid #f0f0f0;padding-bottom:15px;padding-left:15px;padding-right:15px;padding-top:15px;vertical-align:top;border-top:0;border-bottom:0">
                            <table class="divider_block block-1" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0">
                              <tr>
                                <td class="pad">
                                  <div class="alignment" align="center">
                                    <table border="0" cellpadding="0" cellspacing="0" role="presentation" width="100%" style="mso-table-lspace:0;mso-table-rspace:0">
                                      <tr>
                                        <td class="divider_inner" style="font-size:1px;line-height:1px;border-top:1px solid #bbb">
                                          <span>&#8202;</span>
                                        </td>
                                      </tr>
                                    </table>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <table class="text_block block-2" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;word-break:break-word">
                              <tr>
                                <td class="pad">
                                  <div style="font-family:sans-serif">
                                    <div class style="font-size:12px;font-family:Tahoma,Verdana,Segoe,sans-serif;mso-line-height-alt:18px;color:#171717;line-height:1.5">
                                      <p style="margin:0;font-size:12px;mso-line-height-alt:30px">
                                        <span style="font-size:20px;">
                                          <strong>
                                            <span style>কেন এ লগইন কোডটি দরকার?</span>
                                            <br>
                                          </strong>
                                        </span>
                                      </p>
                                    </div>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <table class="text_block block-3" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;word-break:break-word">
                              <tr>
                                <td class="pad">
                                  <div style="font-family:sans-serif">
                                    <div class style="font-size:12px;font-family:Tahoma,Verdana,Segoe,sans-serif;mso-line-height-alt:18px;color:#595959;line-height:1.5">
                                      <p style="margin:0;mso-line-height-alt:27px">
                                        <span style="font-size:18px;">বহুব্রীহি'র নতুন ওয়েবসাইটে আপনার অ্যাকাউন্টটি ট্রান্সফার করা হচ্ছে। ট্রান্সফার সম্পূর্ণ হবার পর থেকে আপনি শুধু মোবাইল ফোন নাম্বার দিয়ে লগইন করতে পারবেন। অ্যাকাউন্ট ট্রান্সফারকে নিরাপদ ও ভেরিফাইড করার জন্য এ লগইন কোডটি দরকার। <br>
                                        </span>
                                      </p>
                                    </div>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <table class="divider_block block-4" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0">
                              <tr>
                                <td class="pad">
                                  <div class="alignment" align="center">
                                    <table border="0" cellpadding="0" cellspacing="0" role="presentation" width="100%" style="mso-table-lspace:0;mso-table-rspace:0">
                                      <tr>
                                        <td class="divider_inner" style="font-size:1px;line-height:1px;border-top:1px solid #bbb">
                                          <span>&#8202;</span>
                                        </td>
                                      </tr>
                                    </table>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <table class="text_block block-5" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;word-break:break-word">
                              <tr>
                                <td class="pad">
                                  <div style="font-family:sans-serif">
                                    <div class style="font-size:12px;font-family:Tahoma,Verdana,Segoe,sans-serif;mso-line-height-alt:18px;color:#171717;line-height:1.5">
                                      <p style="margin:0;font-size:12px;mso-line-height-alt:30px">
                                        <span style="font-size:20px;">
                                          <strong>এখনো কনফিউজড?<br>
                                          </strong>
                                        </span>
                                      </p>
                                    </div>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <table class="text_block block-6" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;word-break:break-word">
                              <tr>
                                <td class="pad">
                                  <div style="font-family:sans-serif">
                                    <div class style="font-size:12px;font-family:Tahoma,Verdana,Segoe,sans-serif;mso-line-height-alt:18px;color:#595959;line-height:1.5">
                                      <p style="margin:0;mso-line-height-alt:27px">
                                        <span style="font-size:18px;">কনফিউশন দূর করতে <a href="#" target="_blank" style="text-decoration: underline; color: #ce268c;" rel="noopener">এ ভিডিওটি দেখুন</a>। </span>
                                      </p>
                                    </div>
                                  </div>
                                </td>
                              </tr>
                            </table>
                            <table class="divider_block block-7" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0">
                              <tr>
                                <td class="pad">
                                  <div class="alignment" align="center">
                                    <table border="0" cellpadding="0" cellspacing="0" role="presentation" width="100%" style="mso-table-lspace:0;mso-table-rspace:0">
                                      <tr>
                                        <td class="divider_inner" style="font-size:1px;line-height:1px;border-top:1px solid #bbb">
                                          <span>&#8202;</span>
                                        </td>
                                      </tr>
                                    </table>
                                  </div>
                                </td>
                              </tr>
                            </table>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </td>
                </tr>
              </tbody>
            </table>
            <table class="row row-5" align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;background-color:#f6f6f6">
              <tbody>
                <tr>
                  <td>
                    <table class="row-content stack" align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;background-color:#fff;border-radius:0;color:#000;width:600px" width="600">
                      <tbody>
                        <tr>
                          <td class="column column-1" width="50%" style="mso-table-lspace:0;mso-table-rspace:0;font-weight:400;text-align:left;padding-bottom:5px;padding-top:5px;vertical-align:top;border-top:0;border-right:0;border-bottom:0;border-left:0">
                            <table class="text_block block-1" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;word-break:break-word">
                              <tr>
                                <td class="pad">
                                  <div style="font-family:sans-serif">
                                    <div class style="font-size:14px;font-family:Tahoma,Verdana,Segoe,sans-serif;mso-line-height-alt:16.8px;color:#1e1e1e;line-height:1.2">
                                      <p style="margin:0;font-size:14px;text-align:center;mso-line-height-alt:16.8px">
                                        <a href="https://bohubrihi.com" target="_blank" style="text-decoration: none; color: #ce268c;" rel="noopener">ওয়েবসাইট ভিজিট করুন</a>
                                      </p>
                                    </div>
                                  </div>
                                </td>
                              </tr>
                            </table>
                          </td>
                          <td class="column column-2" width="50%" style="mso-table-lspace:0;mso-table-rspace:0;font-weight:400;text-align:left;padding-bottom:5px;padding-top:5px;vertical-align:top;border-top:0;border-right:0;border-bottom:0;border-left:0">
                            <table class="text_block block-1" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;word-break:break-word">
                              <tr>
                                <td class="pad">
                                  <div style="font-family:sans-serif">
                                    <div class style="font-size:12px;font-family:Tahoma,Verdana,Segoe,sans-serif;mso-line-height-alt:14.399999999999999px;color:#1e1e1e;line-height:1.2">
                                      <p style="margin:0;font-size:12px;text-align:center;mso-line-height-alt:14.399999999999999px">
                                        <span style="font-size:14px;">কাস্টমার সাপোর্টঃ <a href="tel:+8801916633509" target="_self" title="tel:+8801916633509" style="text-decoration: underline; color: #ce268c;">+8801916-633509</a>
                                        </span>
                                      </p>
                                    </div>
                                  </div>
                                </td>
                              </tr>
                            </table>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </td>
                </tr>
              </tbody>
            </table>
            <table class="row row-6" align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;background-color:#f6f6f6">
              <tbody>
                <tr>
                  <td>
                    <table class="row-content stack" align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;background-color:#ddd;color:#000;width:600px" width="600">
                      <tbody>
                        <tr>
                          <td class="column column-1" width="100%" style="mso-table-lspace:0;mso-table-rspace:0;font-weight:400;text-align:left;padding-bottom:5px;padding-top:5px;vertical-align:top;border-top:0;border-right:0;border-bottom:0;border-left:0">
                            <table class="text_block block-1" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace:0;mso-table-rspace:0;word-break:break-word">
                              <tr>
                                <td class="pad">
                                  <div style="font-family:sans-serif">
                                    <div class style="font-size:12px;font-family:Tahoma,Verdana,Segoe,sans-serif;mso-line-height-alt:14.399999999999999px;color:#1e1e1e;line-height:1.2">
                                      <p style="margin:0;text-align:center;font-size:13px;mso-line-height-alt:15.6px">
                                        <span style="font-size:15px;">
                                          <strong>আমাদের ঠিকানা</strong>
                                        </span>
                                      </p>
                                      <p style="margin:0;text-align:center;font-size:13px;mso-line-height-alt:15.6px">
                                        <span style="font-size:13px;">Rangs Paramount Square,</span>
                                      </p>
                                      <p style="margin:0;text-align:center;font-size:13px;mso-line-height-alt:15.6px">
                                        <span style="font-size:13px;">Floor 11, House # 11, Road # 17, Banani C/A,</span>
                                      </p>
                                      <p style="margin:0;text-align:center;font-size:13px;mso-line-height-alt:15.6px">
                                        <span style="font-size:13px;">Dhaka - 1213, Bangladesh.</span>
                                      </p>
                                    </div>
                                  </div>
                                </td>
                              </tr>
                            </table>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
    <div style="background-color:transparent;">
      <div style="Margin: 0 auto;min-width: 320px;max-width: 500px;overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;background-color: transparent;" class="block-grid ">
        <div style="border-collapse: collapse;display: table;width: 100%;background-color:transparent;">
        </div>
      </div>
    </div>
  </body>
</html>`
	body := fmt.Sprintf(template, otp)

	// print the final email body
	fmt.Println(strings.TrimSpace(body))
}
