/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

directives.directive('simditor', function () {
    var toolbar_simditor = [
        'title',
        'bold',
        'italic',
        'underline',
        'strikethrough',
        '|',
        'fontScale',
        'color',
        'ol',             
        'ul',            
        'blockquote',
        'code',         
        'table',
        '|',
        'link',
        'image',
        'hr',          
        '|',
        'indent',
        'outdent',
        'alignment'
    ];
    return {
        require: "?^ngModel",
        link: function (scope, element, attrs, ngModel) {
            element.append("<div style='height:300px;'></div>");
            var toolbar = scope.$eval(attrs.toolbar) || toolbar_simditor;
            scope.simditor = new Simditor({ 
                textarea: element.children()[0],
                //placeholder: "",
                toolbar: toolbar,
                toolbarFloat: true,
                toolbarFloatOffset: 0,
                toolbarHidden: false,
                defalutImage: "images/image.png, images/image.jpg",
                tabIndent: true,
                upload: {
                    url: '',
                    params: null,
                    fileKey: 'upload_file',
                    connectionCount: 3,
                    leaveConfirm: '正在上传文件, 你确定要离开当前页面么?'
                },
                pasteImage: true,
                cleanPaste: false,
                imageButton: [
                    'upload',
                    'external'
                ],
                allowedTags: [
                    'br',
                    'span',
                    'a',
                    'img',
                    'b',
                    'strong',
                    'i',
                    'strike',
                    'u',
                    'font',
                    'p',
                    'ul',
                    'ol',
                    'li',
                    'blockquote',
                    'pre',
                    'code',
                    'h1',
                    'h2',
                    'h3',
                    'h4',
                    'hr'
                ],
                allowedAttributes: {
                    img: ['src', 'alt', 'width', 'height', 'data-non-image'],
                    a: ['href', 'target'],
                    font: ['color'],
                    code: ['class']
                },
                allowedStyles: {
                    span: ['color', 'font-size'],
                    b: ['color'],
                    i: ['color'],
                    strong: ['color'],
                    strike: ['color'],
                    u: ['color'],
                    p: ['margin-left', 'text-align'],
                    h1: ['margin-left', 'text-align'],
                    h2: ['margin-left', 'text-align'],
                    h3: ['margin-left', 'text-align'],
                    h4: ['margin-left', 'text-align']
                },
                codeLanguages: [
                    { name: 'Bash', value: 'bash' },
                    { name: 'C++', value: 'c++' },
                    { name: 'C#', value: 'cs' },
                    { name: 'CSS', value: 'css' },
                    { name: 'Erlang', value: 'erlang' },
                    { name: 'Less', value: 'less' },
                    { name: 'Sass', value: 'sass' },
                    { name: 'Diff', value: 'diff' },
                    { name: 'CoffeeScript', value: 'coffeescript' },
                    { name: 'HTML,XML', value: 'html' },
                    { name: 'JSON', value: 'json' },
                    { name: 'Java', value: 'java' },
                    { name: 'JavaScript', value: 'js' },
                    { name: 'Markdown', value: 'markdown' },
                    { name: 'Objective C', value: 'oc' },
                    { name: 'PHP', value: 'php' },
                    { name: 'Perl', value: 'parl' },
                    { name: 'Python', value: 'python' },
                    { name: 'Ruby', value: 'ruby' },
                    { name: 'SQL', value: 'sql'}
                ],
                params: {}
            });

            ngModel.$render = function () {
                var $target = element.find('.simditor-body');
                scope.simditor.focus();
                $target.html(ngModel.$viewValue);
            };

            scope.simditor.on('valuechanged', function () {
                scope.$apply(function(){
                    var $target = element.find('.simditor-body');
                    ngModel.$setViewValue($target.html());

                    if (attrs.ngRequired != undefined && attrs.ngRequired != "false") {
                        var text = $target.text();

                        if(text.trim() === "") {
                            ngModel.$setValidity("required", false);
                        } else {
                            ngModel.$setValidity("required", true);
                        }
                    }
                });
            });
        }
    };
});
