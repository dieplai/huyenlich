$ErrorActionPreference = "Stop"

$tarotDir = Join-Path (Get-Location) "data\tarot"
$deckPath = Join-Path $tarotDir "rider-waite-smith.json"
$deck = Get-Content -Raw -Encoding UTF8 -LiteralPath $deckPath | ConvertFrom-Json

$suitProfiles = @{
    wands = @{
        name_vi = "Gậy"
        element_vi = "Hỏa"
        field_vi = "hành động, đam mê, động lực và hướng tiến lên"
        love_vi = "nhiệt, hấp dẫn, chủ động, ham muốn và cách hai người tạo chuyển động"
        career_vi = "sáng kiến, cạnh tranh, tốc độ, dự án mới và tinh thần dẫn dắt"
        money_vi = "cơ hội đến từ hành động, kinh doanh, thử nghiệm và quyết định nhanh"
        inner_vi = "lửa bên trong: điều khiến bạn muốn sống, muốn thử, muốn bước tới"
        shadow_vi = "nóng vội, thiếu kiên nhẫn, phản ứng theo hứng hoặc đốt năng lượng quá nhanh"
        advice_vi = "chọn một hành động có hướng, đừng để cảm hứng tan thành bốc đồng"
        warning_vi = "đam mê mạnh cần được kiểm chứng bằng sự nhất quán"
    }
    cups = @{
        name_vi = "Cốc"
        element_vi = "Thủy"
        field_vi = "cảm xúc, tình yêu, trực giác và sự kết nối"
        love_vi = "cảm xúc, sự mở lòng, nhu cầu được yêu và cách hai bên trao nhận"
        career_vi = "môi trường cảm xúc, quan hệ đồng đội, sáng tạo và công việc cần đồng cảm"
        money_vi = "quyết định tài chính chịu ảnh hưởng bởi cảm xúc, giá trị cá nhân và cảm giác an toàn"
        inner_vi = "dòng cảm xúc bên trong, điều bạn khao khát và cách bạn tự chăm mình"
        shadow_vi = "mơ hồ, phụ thuộc cảm xúc, lý tưởng hóa hoặc ngập trong cảm giác"
        advice_vi = "gọi đúng tên cảm xúc rồi mới quyết định"
        warning_vi = "cảm xúc thật cần hành động thật để không thành mơ mộng"
    }
    swords = @{
        name_vi = "Kiếm"
        element_vi = "Khí"
        field_vi = "suy nghĩ, sự thật, giao tiếp và xung đột tinh thần"
        love_vi = "lời nói, hiểu lầm, ranh giới tinh thần và sự thật cần được nói ra"
        career_vi = "quyết định, chiến lược, tranh luận, dữ kiện và áp lực trí óc"
        money_vi = "phân tích rủi ro, hợp đồng, quyết định dựa trên dữ kiện và tránh suy nghĩ cực đoan"
        inner_vi = "tiếng nói trong đầu, nỗi sợ, niềm tin và cách bạn diễn giải sự việc"
        shadow_vi = "lo âu, lạnh cảm xúc, phán xét, phòng thủ hoặc nói sắc quá"
        advice_vi = "nói rõ, kiểm chứng dữ kiện và cắt bớt suy diễn"
        warning_vi = "sự thật không nên được dùng như vũ khí"
    }
    pentacles = @{
        name_vi = "Tiền"
        element_vi = "Thổ"
        field_vi = "công việc, tiền bạc, thân thể, thói quen và nền tảng thực tế"
        love_vi = "sự ổn định, chăm sóc bằng hành động, thời gian và mức độ đáng tin"
        career_vi = "tay nghề, quy trình, thành quả vật chất, trách nhiệm và sự bền bỉ"
        money_vi = "nguồn lực, thu chi, tích lũy, đầu tư thực tế và cảm giác an toàn"
        inner_vi = "cảm giác đủ, an toàn, giá trị bản thân và mối liên hệ với thân thể"
        shadow_vi = "bám víu, chậm đổi, quá thực dụng hoặc đánh đồng giá trị với vật chất"
        advice_vi = "đưa cảm xúc và ý tưởng xuống hành động cụ thể có thể đo được"
        warning_vi = "ổn định không nên trở thành lý do để mắc kẹt"
    }
}

$rankProfiles = @{
    ace = @{
        stage_vi = "hạt giống mới"
        core_vi = "một cơ hội nguyên sơ vừa mở ra"
        upright_vi = "có tín hiệu mới đáng chú ý, nhưng nó mới là hạt giống và cần được nuôi đúng cách"
        reversed_vi = "cơ hội có đó nhưng bị chậm, bị bỏ lỡ hoặc chưa đủ nền để nảy mầm"
        action_vi = "đón cơ hội bằng một bước nhỏ, không vội biến nó thành cam kết lớn"
        outcome_vi = "một cửa mới mở ra, kết quả phụ thuộc vào cách bạn chăm nó"
        challenge_vi = "khó phân biệt tiềm năng thật với cảm hứng thoáng qua"
        relationship_default_vi = "có tín hiệu mới hoặc mong muốn làm mới, nhưng cần xem nó có được nuôi bằng hành động không"
    }
    two = @{
        stage_vi = "lựa chọn và cân bằng"
        core_vi = "hai lực, hai hướng hoặc hai người đang cần tìm điểm cân bằng"
        upright_vi = "tình huống cần đối thoại, cân bằng hoặc lựa chọn rõ hơn"
        reversed_vi = "sự cân bằng lệch, do dự hoặc né quyết định đang làm tình huống đứng yên"
        action_vi = "đặt hai lựa chọn lên bàn và xem điều gì thật sự đáp lại bạn"
        outcome_vi = "xu hướng phụ thuộc vào việc hai phía có tìm được điểm chung không"
        challenge_vi = "do dự, chờ người khác quyết hộ hoặc giữ hòa khí bằng im lặng"
        relationship_default_vi = "dấu hiệu cần xem là sự đáp lại có hai chiều không"
    }
    three = @{
        stage_vi = "phát triển và biểu đạt"
        core_vi = "năng lượng bắt đầu mở rộng qua giao tiếp, hợp tác hoặc biểu hiện ra ngoài"
        upright_vi = "có tiến triển khi mọi thứ được chia sẻ, thể hiện hoặc cùng xây"
        reversed_vi = "sự phát triển bị chậm vì thiếu phối hợp, thiếu rõ ràng hoặc biểu đạt chưa thật"
        action_vi = "đưa điều đang nghĩ/cảm ra thành trao đổi hoặc kế hoạch chung"
        outcome_vi = "tình huống có thể mở rộng nếu có sự phối hợp thật"
        challenge_vi = "nói nhiều nhưng chưa cùng xây, hoặc có người thứ ba/yếu tố ngoài làm nhiễu"
        relationship_default_vi = "dấu hiệu cần xem là hai bên có cùng góp phần phát triển kết nối không"
    }
    four = @{
        stage_vi = "cấu trúc và ổn định"
        core_vi = "nhu cầu về nền tảng, ranh giới và cảm giác an toàn"
        upright_vi = "tình huống cần được đặt vào cấu trúc rõ hơn để ổn định"
        reversed_vi = "sự ổn định có thể thành cứng, hoặc nền tảng chưa đủ chắc"
        action_vi = "làm rõ ranh giới, nhịp độ và điều kiện tối thiểu"
        outcome_vi = "có thể ổn định nếu cấu trúc lành mạnh; nếu không, sự tù túng sẽ lộ ra"
        challenge_vi = "bám vào an toàn cũ hoặc thiếu nền để đi xa"
        relationship_default_vi = "dấu hiệu cần xem là kết nối này đem lại an toàn hay chỉ giữ bạn trong khuôn quen"
    }
    five = @{
        stage_vi = "biến động và thử thách"
        core_vi = "một điểm căng xuất hiện để buộc tình huống phải điều chỉnh"
        upright_vi = "có xung đột, thiếu hụt hoặc biến động cần nhìn thẳng"
        reversed_vi = "xung đột đang lắng xuống hoặc bị né tránh quá lâu"
        action_vi = "xác định đúng điểm đau, không cố thắng mọi cuộc đấu"
        outcome_vi = "nếu không điều chỉnh, biến động tiếp tục; nếu học được bài, tình huống đổi hướng"
        challenge_vi = "phản ứng vì tổn thương, cạnh tranh hoặc cảm giác thiếu"
        relationship_default_vi = "dấu hiệu cần xem là xung đột này giúp rõ hơn hay chỉ làm cả hai cạn đi"
    }
    six = @{
        stage_vi = "điều hòa và phục hồi"
        core_vi = "năng lượng tìm lại sự cân bằng sau biến động"
        upright_vi = "có cơ hội hỗ trợ, chữa lành hoặc đưa tình huống về nhịp dễ thở hơn"
        reversed_vi = "sự trao nhận lệch hoặc quá khứ vẫn chưa được đặt đúng chỗ"
        action_vi = "đưa mọi thứ về công bằng, hỗ trợ và nhịp lành"
        outcome_vi = "có thể mềm lại nếu hai bên biết trao nhận đúng mức"
        challenge_vi = "một bên cho quá nhiều, hoặc quá khứ kéo hiện tại lệch đi"
        relationship_default_vi = "dấu hiệu cần xem là sự cho nhận có cân bằng và làm bạn nhẹ hơn không"
    }
    seven = @{
        stage_vi = "kiểm tra và chọn lọc"
        core_vi = "một bài kiểm tra về niềm tin, lựa chọn, phòng thủ hoặc sự rõ ràng bên trong"
        upright_vi = "tình huống cần quan sát kỹ và chọn lọc thay vì phản ứng ngay"
        reversed_vi = "bạn có thể đang phòng thủ quá mức, mất kiên nhẫn hoặc né bài học chính"
        action_vi = "chậm lại để xem điều gì thật sự đáng giữ"
        outcome_vi = "xu hướng phụ thuộc vào khả năng nhìn rõ bài kiểm tra đang diễn ra"
        challenge_vi = "nghi ngờ, phòng thủ, mơ hồ hoặc muốn kết quả nhanh"
        relationship_default_vi = "dấu hiệu cần xem là thử thách này làm kết nối trưởng thành hay làm bạn luôn trong thế phòng thủ"
    }
    eight = @{
        stage_vi = "chuyển động và làm chủ"
        core_vi = "năng lượng đi vào quá trình vận hành, rèn luyện hoặc tăng tốc"
        upright_vi = "tình huống có thể tiến triển nếu có kỷ luật và hướng rõ"
        reversed_vi = "tiến trình bị chặn, quá nhanh, quá máy móc hoặc thiếu điều chỉnh"
        action_vi = "đưa năng lượng vào quy trình cụ thể"
        outcome_vi = "có chuyển động rõ nếu không để nó lệch nhịp"
        challenge_vi = "mất kiểm soát tốc độ, thiếu kiên trì hoặc mắc trong lặp lại"
        relationship_default_vi = "dấu hiệu cần xem là hai bên có cùng tạo chuyển động đều không"
    }
    nine = @{
        stage_vi = "kết tinh và tự chủ"
        core_vi = "một giai đoạn gần hoàn tất, đòi hỏi nhìn sâu và giữ năng lượng của mình"
        upright_vi = "bạn đã đi qua nhiều điều và đang chạm tới kết quả hoặc nhận thức quan trọng"
        reversed_vi = "sự gần hoàn tất bị cản bởi mệt mỏi, cô lập hoặc chưa dám nhận bài học"
        action_vi = "ghi nhận điều đã đi qua và bảo vệ phần năng lượng còn lại"
        outcome_vi = "kết quả gần hiện ra, nhưng cần đi nốt đoạn cuối với sự tỉnh táo"
        challenge_vi = "cạn sức trước vạch cuối hoặc giữ mọi thứ một mình quá lâu"
        relationship_default_vi = "dấu hiệu cần xem là kết nối này làm bạn trưởng thành hay chỉ khiến bạn phải tự chịu đựng"
    }
    ten = @{
        stage_vi = "hoàn tất và chuyển pha"
        core_vi = "một chu kỳ đi tới điểm đầy hoặc điểm kết thúc"
        upright_vi = "tình huống đang đạt đỉnh, hoàn tất hoặc cần chuyển sang vòng mới"
        reversed_vi = "chu kỳ chưa được khép đúng cách hoặc bạn đang kéo dài phần đã quá tải"
        action_vi = "khép vòng, dọn phần dư và chuẩn bị chuyển pha"
        outcome_vi = "một chương có xu hướng kết thúc để mở chương khác"
        challenge_vi = "bám vào chu kỳ cũ vì sợ khoảng trống sau đó"
        relationship_default_vi = "dấu hiệu cần xem là hai người đang đi tới cột mốc trưởng thành hơn hay tới điểm quá tải"
    }
    page = @{
        stage_vi = "người học và tín hiệu mới"
        core_vi = "một năng lượng non mới, tò mò hoặc một thông điệp đang xuất hiện"
        upright_vi = "có tin mới, sự tò mò hoặc cơ hội học cách tiếp cận khác"
        reversed_vi = "năng lượng còn non, thiếu nhất quán hoặc chưa biết diễn đạt đúng"
        action_vi = "hỏi, học, thử nhỏ và đừng vội đòi sự trưởng thành hoàn chỉnh"
        outcome_vi = "xu hướng là một tín hiệu/tin nhắn/bước học mới hơn là kết luận cuối"
        challenge_vi = "thiếu kinh nghiệm, nói chưa tới hoặc hứng lên rồi thôi"
        relationship_default_vi = "dấu hiệu cần xem là sự quan tâm có trưởng thành dần không hay chỉ là tò mò nhất thời"
    }
    knight = @{
        stage_vi = "người hành động và xung lực"
        core_vi = "năng lượng tiến tới, theo đuổi hoặc đẩy tình huống sang trạng thái động"
        upright_vi = "có hành động, động lực hoặc một người/sự kiện đang lao tới"
        reversed_vi = "hành động quá đà, thiếu ổn định hoặc bị chặn giữa đường"
        action_vi = "hành động nhưng phải biết mình đang theo đuổi điều gì"
        outcome_vi = "tình huống có thể chuyển động nhanh, tốt hay xấu tùy mức định hướng"
        challenge_vi = "bốc đồng, cực đoan hoặc thiếu cam kết với đoạn đường dài"
        relationship_default_vi = "dấu hiệu cần xem là người kia tiến tới vì rõ lòng hay chỉ bị cảm xúc/ham muốn kéo"
    }
    queen = @{
        stage_vi = "người làm chủ bên trong"
        core_vi = "năng lượng trưởng thành, tiếp nhận, nuôi dưỡng và làm chủ suit ở tầng nội tâm"
        upright_vi = "có sự trưởng thành cảm xúc/nội lực và khả năng giữ không gian cho điều quan trọng"
        reversed_vi = "năng lượng chăm sóc/làm chủ bị lệch, quá tải hoặc thiếu ranh giới"
        action_vi = "giữ phẩm chất của mình nhưng đặt ranh giới mềm"
        outcome_vi = "tình huống có thể dịu và sâu hơn nếu có sự trưởng thành bên trong"
        challenge_vi = "ôm quá nhiều, phản ứng từ tổn thương hoặc thiếu tự chăm"
        relationship_default_vi = "dấu hiệu cần xem là tình cảm có được nuôi bằng sự trưởng thành hay bằng việc một người ôm hết"
    }
    king = @{
        stage_vi = "người quản trị và chịu trách nhiệm"
        core_vi = "năng lượng trưởng thành, điều phối, thể hiện ra ngoài và chịu trách nhiệm với suit"
        upright_vi = "cần sự rõ ràng, trách nhiệm và khả năng dẫn dắt năng lượng này một cách chín chắn"
        reversed_vi = "quyền lực/lãnh đạo của suit bị lệch: kiểm soát, lạnh, bốc đồng hoặc vật chất hóa quá mức tùy suit"
        action_vi = "hành động như người có trách nhiệm, không chỉ như người có cảm xúc"
        outcome_vi = "kết quả có thể ổn nếu có trách nhiệm và sự làm chủ rõ"
        challenge_vi = "lạm quyền, kiểm soát hoặc thiếu trưởng thành trong cách dùng năng lượng"
        relationship_default_vi = "dấu hiệu cần xem là người kia có đủ trưởng thành để chịu trách nhiệm với điều họ tạo ra không"
    }
}

$rankNameVi = @{
    ace = "Át"; two = "Hai"; three = "Ba"; four = "Bốn"; five = "Năm"; six = "Sáu"; seven = "Bảy"; eight = "Tám"; nine = "Chín"; ten = "Mười"; page = "Tiểu Đồng"; knight = "Kỵ Sĩ"; queen = "Nữ Hoàng"; king = "Vua"
}

$rankToKey = @{
    ace = "ace"; two = "2"; three = "3"; four = "4"; five = "5"; six = "6"; seven = "7"; eight = "8"; nine = "9"; ten = "10"; page = "page"; knight = "knight"; queen = "queen"; king = "king"
}

$sourceBasis = @("waite", "golden_dawn_book_t", "joan_bunning_learning_tarot", "benebell_holistic_tarot", "product_editorial")

foreach ($suit in @("wands", "cups", "swords", "pentacles")) {
    $profile = $suitProfiles[$suit]
    $minorCards = @($deck.cards | Where-Object { $_.suit -eq $suit } | Sort-Object deck_index)
    $cards = @()

    foreach ($baseCard in $minorCards) {
        $rank = [string]$baseCard.rank
        $rankProfile = $rankProfiles[$rank]
        $rankLabel = $rankNameVi[$rank]
        $numberKey = $rankToKey[$rank]

        $relationshipSignals = [ordered]@{
            default = "$($rankProfile.relationship_default_vi). Với $($profile.name_vi), hãy quan sát $($profile.love_vi)."
            existing_relationship = "Nếu hỏi có nên tiếp tục, $($baseCard.name_vi) yêu cầu xem $($profile.love_vi) đang giúp hai bên tiến gần hơn hay chỉ lặp lại cùng mô thức. Trục chính là $($rankProfile.stage_vi)."
            situationship = "Trong mập mờ, lá này nhấn vào $($rankProfile.core_vi) trong vùng $($profile.field_vi). Dấu hiệu cần xem là có chuyển thành hành động rõ hay chỉ dừng ở cảm giác."
            ex = "Với người cũ, $($baseCard.name_vi) hỏi liệu bài học $($rankProfile.stage_vi) đã được hiểu chưa, hay hai bên chỉ đang chạm lại vì thói quen cũ."
            no_contact = "Nếu đang im lặng, lá này cho thấy năng lượng $($profile.name_vi) đang bị giữ lại. Hãy xem dấu hiệu liên hệ nếu có mang tính trưởng thành hay chỉ là xung lực."
        }

        $card = [ordered]@{
            id = $baseCard.id
            source_basis = $sourceBasis
            core = [ordered]@{
                essence_vi = "$($baseCard.name_vi) kết hợp năng lượng $($rankProfile.stage_vi) với lĩnh vực $($profile.field_vi). Cốt lõi là $($rankProfile.core_vi)."
                keywords_vi = @($rankProfile.stage_vi, $profile.field_vi, $rankProfile.core_vi)
                shadow_keywords_vi = @($profile.shadow_vi, $rankProfile.challenge_vi)
            }
            orientation = [ordered]@{
                upright = [ordered]@{
                    general_vi = "$($rankProfile.upright_vi). Với suit $($profile.name_vi), trọng tâm rơi vào $($profile.field_vi)."
                    love_vi = "Trong tình yêu, $($baseCard.name_vi) nói về $($profile.love_vi). $($rankProfile.upright_vi)."
                    career_vi = "Về công việc, lá này đặt $($rankProfile.stage_vi) vào $($profile.career_vi). $($rankProfile.upright_vi)."
                    money_vi = "Về tiền bạc, lá này liên quan $($profile.money_vi). $($rankProfile.upright_vi)."
                    inner_work_vi = "Ở tầng nội tâm, lá này chạm tới $($profile.inner_vi). $($rankProfile.upright_vi)."
                    advice_vi = "$($rankProfile.action_vi). $($profile.advice_vi)."
                    warning_vi = "$($profile.warning_vi). Điểm cần tránh là $($rankProfile.challenge_vi)."
                }
                reversed = [ordered]@{
                    general_vi = "$($rankProfile.reversed_vi). Mặt lệch của suit $($profile.name_vi) là $($profile.shadow_vi)."
                    love_vi = "Trong tình yêu, lá ngược cho thấy $($profile.love_vi) đang bị chặn hoặc lệch. $($rankProfile.reversed_vi)."
                    career_vi = "Về công việc, lá ngược cảnh báo $($profile.career_vi) đang thiếu nhịp, thiếu hướng hoặc thiếu sự trưởng thành phù hợp. $($rankProfile.reversed_vi)."
                    money_vi = "Về tiền bạc, cần cẩn trọng vì $($profile.money_vi) có thể bị chi phối bởi mặt lệch: $($profile.shadow_vi)."
                    inner_work_vi = "Ở tầng nội tâm, lá ngược cho thấy bạn cần nhìn lại $($profile.inner_vi) khi nó bị kéo bởi $($profile.shadow_vi)."
                    advice_vi = "$($rankProfile.action_vi), nhưng trước hết hãy xử lý phần bị chặn: $($rankProfile.challenge_vi)."
                    warning_vi = "Nếu không điều chỉnh, $($profile.shadow_vi) sẽ làm $($rankProfile.stage_vi) biến thành điểm nghẽn."
                }
            }
            position_overrides = [ordered]@{
                situation = "Tình huống đang biểu hiện $($rankProfile.stage_vi) trong vùng $($profile.field_vi)."
                action = $rankProfile.action_vi
                outcome = $rankProfile.outcome_vi
                challenge = $rankProfile.challenge_vi
                future = $rankProfile.outcome_vi
            }
            relationship_signals = $relationshipSignals
            symbols = @(
                [ordered]@{
                    symbol = "suit $($profile.name_vi)"
                    meaning_vi = "Đưa trọng tâm về $($profile.field_vi), thuộc nguyên tố $($profile.element_vi)."
                },
                [ordered]@{
                    symbol = "rank $rankLabel"
                    meaning_vi = "Đưa câu chuyện vào giai đoạn $($rankProfile.stage_vi)."
                }
            )
            reader_notes_vi = @(
                "Đọc lá này bằng cách kết hợp rank trước, suit sau: $($rankProfile.stage_vi) + $($profile.field_vi).",
                "Không đổi nghĩa gốc của lá; V2 chỉ giúp diễn giải theo ngữ cảnh câu hỏi và position."
            )
        }

        $cards += $card
    }

    $out = [ordered]@{
        schema_version = "2.0-minor-$suit"
        deck_id = "rider-waite-smith"
        coverage = [ordered]@{
            card_count = $cards.Count
            notes_vi = "V2 cho 14 lá suit $($profile.name_vi), sinh từ rank framework + suit framework đã lưu trong rule data."
        }
        cards = $cards
    }

    $outPath = Join-Path $tarotDir "rider-waite-smith.v2.$suit.json"
    $json = $out | ConvertTo-Json -Depth 30
    [System.IO.File]::WriteAllText($outPath, $json, [System.Text.UTF8Encoding]::new($false))
    Write-Host "Wrote $outPath cards=$($cards.Count)"
}
