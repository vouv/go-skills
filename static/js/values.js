'use strict';

angular.module('tour.values', []).

// List of modules with description and lessons in it.
value('tableOfContents', [{
    'id': 'welcome',
    'title': 'Getting Started',
    'description': '面向代码的计算机基础知识汇总 for Gopher，用于知识巩固、学习提高或技术面试。',
    'lessons': ['welcome']
}, {
    'id': 'core',
    'title': 'Go 语言核心',
    'description': '',
    'lessons': ['core-type', 'core-concurrent', 'core-reflect', 'core-memory', 'core-gc', 'core-sdk', 'core-trap']
}, {
    'id': 'algorithm',
    'title': '数据结构与算法',
    'description': '',
    'lessons': ['algo-before', 'algo-basic', 'algo-advanced','algo-mind', 'algo-sorts', 'algo-string']
}, {
    'id': 'network',
    'title': '计算机网络',
    'description': '',
    'lessons': ['network-basic', 'network-network-layer', 'network-transport-layer', 'network-application-layer']
}, {
    'id': 'operating-system',
    'title': '操作系统',
    'description': '',
    'lessons': ['os-process', 'os-socket', 'os-file-system']
}, {
    'id': 'database',
    'title': '数据库',
    'description': '',
    'lessons': ['db-basic', 'db-sql', 'db-mysql', 'db-redis']
},  {
    'id': 'distributed',
    'title': '分布式',
    'description': '',
    'lessons': ['distributed-basic', 'distributed-storage', 'distributed-rpc']
}, {
    'id': 'structure',
    'title': '架构思维',
    'description': '',
    'lessons': ['structure']
}, {
    'id': 'practice',
    'title': '技术修炼',
    'description': '',
    'lessons': ['practice']
}, {
    'id': 'interview',
    'title': '面试题解析',
    'description': '',
    'lessons': ['interview-offer', 'interview-code']
}]).

// translation
value('translation', {
    'off': '关',
    'on': '开',
    'syntax': '语法高亮',
    'lineno': '显示行号',
    'reset': 'Reset Slide',
    'format': '格式化源代码',
    'kill': 'Kill Program',
    'run': 'Run',
    'compile': 'Compile and Run',
    'more': 'Options',
    'toc': '显示目录',
    'prev': '上一页',
    'next': '下一页',
    'waiting': 'Program Executing...',
    'formatting': 'Fmt Executing...',
    'errcomm': '远端服务器通信错误.',
    'submit-feedback': '发送页面反馈',
    'github-follow': '前往Github',

    // GitHub issue template: update repo and messaging when translating.
    'github-repo': 'github.com/vouv/go-skills',
    'issue-title': '[BUG|建议] 标题',
    'issue-message': '',
    'context': 'content',
}).

// Config for codemirror plugin
value('ui.config', {
    codemirror: {
        mode: 'text/x-go',
        matchBrackets: true,
        lineNumbers: true,
        autofocus: true,
        indentWithTabs: true,
        indentUnit: 4,
        tabSize: 4,
        lineWrapping: true,
        extraKeys: {
            'Shift-Enter': function() {
                $('#run').click();
            },
            'Ctrl-Enter': function() {
                $('#format').click();
            },
            'PageDown': function() {
                return false;
            },
            'PageUp': function() {
                return false;
            },
        },
        // TODO: is there a better way to do this?
        // AngularJS values can't depend on factories.
        onChange: function() {
            if (window.codeChanged !== null) window.codeChanged();
        }
    }
});

// .value('mapping', {});