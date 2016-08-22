import Ember from "ember";

export default Ember.Helper.helper(function(inputs, options) {

    let value = inputs[0];

    var output = '\
        a, a:hover, a:visited { \
            color:rgb('+value+'); \
        } \
        .colour-main { \
            color:rgb('+value+') !important; \
        } \
        .OnlinePlatforms-frame { \
            background-color:rgb('+value+'); \
        } \
        .OnlinePlatforms-system-sur-menu { \
            background-color:rgb('+value+'); \
        } \
        .button-special-main { \
            background-color:rgba('+value+',0.8); \
            border-color:rgba(0,0,0,0.1); \
        } \
        .button-special-main:hover { \
            background-color:rgb('+value+'); \
        } \
        .OnlinePlatforms-tabs-menu-tab.selected { \
            color:rgb('+value+'); \
            background-color:rgba('+value+',0.05); \
        } \
        .OnlinePlatforms-dashboard-inner > div.row > div .head { \
            color:rgb('+value+'); \
        } \
        op-list-view op-list-item.selected { \
            background-color:rgba('+value+',0.2); \
        } \
        op-grid-view op-grid-item.selected text { \
            color:rgb('+value+'); \
        } \
        op-grid-view op-grid-item.selected icon { \
            border-color:rgb('+value+'); \
        } \
        trigger-plan trigger-plan-node.selected { \
            border-color:rgb('+value+'); \
        } \
        t.checked { \
            background-color:rgb('+value+'); \
        } \
        c.checked { \
            box-shadow: inset 0 0 0 15px rgb('+value+'); \
        } \
        op-input-select op-input-select-text[disabled=false]:hover { \
            color:rgb('+value+'); \
        } \
        op-input-select op-input-select-menu op-input-option:hover { \
            background-color:rgb('+value+'); \
        } \
        op-input-select op-input-select-menu op-input-option.isset:hover { \
            background-color:rgb('+value+'); \
        } \
        op-input-autocomplete-results-result:hover { \
            background-color:rgb('+value+'); \
        } \
        input[type=text]:focus, \
        input[type=password]:focus, \
        input[type=email]:focus, \
        input[type=phone]:focus, \
        input[type=tel]:focus, \
        input[type=url]:focus, \
        input[type=address]:focus, \
        input[type=text]:focus, \
        input[type=number]:focus, \
        textarea:focus \
        { \
            box-shadow: 0 2px 0 rgb('+value+'); \
        } \
    ';

    return new Ember.Handlebars.SafeString( '<style>' + output + '</style>' );

});