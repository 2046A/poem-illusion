/**
 * Created by ivan on 2016/8/11.
 */

$('.tabular.menu .item').tab();

$('.message .close')
    .on('click', function() {
        $(this)
            .closest('.message')
            .transition('fade')
        ;
    })
;
