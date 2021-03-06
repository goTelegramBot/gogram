package types

import (
    "io"
    "os"
)

type StickerSet struct {
    Name          string    `json:"name"`
    Title         string    `json:"title"`
    IsAnimated    bool      `json:"is_animated"`
    ContainsMasks bool      `json:"contains_masks"`
    Stickers      []Sticker `json:"stickers"`
    Thumb         PhotoSize `json:"thumb"`
}
type InputFile struct {
    FileID     string
    FileURL    string
    FilePath   string
    FileReader io.Reader
    FileBytes  []byte
}

func (file InputFile) OnDisk() bool {
    _, err := os.Stat(file.FilePath)
    if os.IsNotExist(err) {
        return false
    }
    return true
}

type File struct {
    FileID       string `json:"file_id"`
    FileUniqueID string `json:"file_unique_id"`
    FileSize     int    `json:"file_size"`
    FilePath     string `json:"file_path"`
}

type UserProfilePhotos struct {
    TotalCount int           `json:"total_count"`
    Photos     [][]PhotoSize `json:"photos"`
}
type BotCommand struct {
    Text        string `json:"command"`
    Description string `json:"description"`
}

type ResponseParameters struct {
    MigrateToChatID string `json:"migrate_to_chat_id"`
    RetryAfter      string `json:"retry_after"`
}

type ReplyMarkup struct {
    InlineKeyboardMarkup *InlineKeyboardReplyMarkup `json:"inline_keyboard,omitempty"`

    ReplyKeyboardMarkup *KeyboardMarkup `json:"keyboard,omitempty"`

    ForceReply          bool `json:"force_reply,omitempty"`
    ReplyKeyboardRemove bool `json:"remove_keyboard,omitempty"`
    Selective           bool `json:"selective,omitempty"`
}
type InlineKeyboardReplyMarkup struct {
    InlineKeyboardMarkup [][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}
type User struct {
    ID                      int    `json:"id,omitempty"`
    IsBot                   bool   `json:"is_bot,omitempty"`
    FirstName               string `json:"first_name,omitempty"`
    LastName                string `json:"last_name,omitempty"`
    UserName                string `json:"username,omitempty"`
    LanguageCode            string `json:"language_code,omitempty"`
    CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
    CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
    SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
}

type MessageEntity struct {
    Type     string `json:"type,omitempty"`
    Offset   int64  `json:"offset,omitempty"`
    Length   int64  `json:"length,omitempty"`
    Url      string `json:"url,omitempty"`
    User     *User  `json:"user,omitempty"`
    Language string `json:"language,omitempty"`
}

type InlineKeyboardButton struct {
    Text string `json:"text,omitempty"`
    Url  string `json:"url,omitempty"`

    LoginUrl *LoginUrl `json:"login_url,omitempty"`

    CallbackData string `json:"callback_data,omitempty"`

    SwitchInlineQuery string `json:"switch_inline_query,omitempty"`

    SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`

    CallbackGame *CallbackGame `json:"callback_game,omitempty"`

    Pay bool `json:"pay,omitempty"`
}
type LoginUrl struct {
    Url                string `json:"url,omitempty"`
    ForwardText        string `json:"forward_text,omitempty"`
    BotUsername        string `json:"bot_username,omitempty"`
    RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

type KeyboardMarkup struct {
    Keyboard        [][]KeyboardButton `json:"keyboard,omitempty"`
    ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
    OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
    Selective       bool               `json:"selective,omitempty"`
}

type CallbackGame interface{}

type KeyboardButton struct {
    Text            string                  `json:"text,omitempty"`
    RequestContact  bool                    `json:"request_contact,omitempty"`
    RequestLocation bool                    `json:"request_location,omitempty"`
    RequestPoll     *KeyboardButtonPollType `json:"request_poll,omitempty"`
}

type KeyboardButtonPollType struct {
    Type string `json:"type,omitempty"`
}

type Update struct {
    UpdateID           int                 `json:"update_id,omitempty"`
    Message            *Message            `json:"message,omitempty"`
    EditedMessage      *Message            `json:"edited_message,omitempty"`
    ChannelPost        *Message            `json:"channel_post,omitempty"`
    EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
    InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
    ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
    CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
    ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
    PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
    Poll               *Poll               `json:"poll,omitempty"`
    PollAnswer         *PollAnswer         `json:"poll_answer,omitempty"`
    MyChatMember       *ChatMemberUpdated  `json:"my_chat_member,omitempty"`
    ChatMember         *ChatMemberUpdated  `json:"chat_member,omitempty"`
}

type Message struct {
    MessageId                     int                            `json:"message_id,omitempty"`
    From                          *User                          `json:"from,omitempty"`
    SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
    Date                          int64                          `json:"date,omitempty"`
    Chat                          Chat                           `json:"chat,omitempty"`
    ForwardFrom                   *User                          `json:"forward_from,omitempty"`
    ForwardFromChat               *Chat                          `json:"forward_from_chat,omitempty"`
    ForwardFromMessageId          int64                          `json:"forward_from_message_id,omitempty"`
    ForwardSignature              string                         `json:"forward_signature,omitempty"`
    ForwardSenderName             string                         `json:"forward_sender_name,omitempty"`
    ForwardDate                   int64                          `json:"forward_date,omitempty"`
    ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
    ViaBot                        *User                          `json:"via_bot,omitempty"`
    EditDate                      int64                          `json:"edit_date,omitempty"`
    MediaGroupId                  string                         `json:"media_group_id,omitempty"`
    AuthorSignature               string                         `json:"author_signature,omitempty"`
    Text                          string                         `json:"text,omitempty"`
    Entities                      []MessageEntity                `json:"entities,omitempty"`
    Animation                     *Animation                     `json:"animation,omitempty"`
    Audio                         *Audio                         `json:"audio,omitempty"`
    Document                      *Document                      `json:"document,omitempty"`
    Photo                         []PhotoSize                    `json:"photo,omitempty"`
    Sticker                       *Sticker                       `json:"sticker,omitempty"`
    Video                         *Video                         `json:"video,omitempty"`
    VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
    Voice                         *Voice                         `json:"voice,omitempty"`
    Caption                       string                         `json:"caption,omitempty"`
    CaptionEntities               []MessageEntity                `json:"caption_entities,omitempty"`
    Contact                       *Contact                       `json:"contact,omitempty"`
    Dice                          *Dice                          `json:"dice,omitempty"`
    Game                          *Game                          `json:"game,omitempty"`
    Poll                          *Poll                          `json:"poll,omitempty"`
    Venue                         *Venue                         `json:"venue,omitempty"`
    Location                      *Location                      `json:"location,omitempty"`
    NewChatMembers                []User                         `json:"new_chat_members,omitempty"`
    LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
    NewChatTitle                  string                         `json:"new_chat_title,omitempty"`
    NewChatPhoto                  []PhotoSize                    `json:"new_chat_photo,omitempty"`
    DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`
    GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`
    SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`
    ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`
    MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
    MigrateToChatId               int64                          `json:"migrate_to_chat_id,omitempty"`
    MigrateFromChatId             int64                          `json:"migrate_from_chat_id,omitempty"`
    PinnedMessage                 *Message                       `json:"pinned_message,omitempty"`
    Invoice                       *Invoice                       `json:"invoice,omitempty"`
    SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
    ConnectedWebsite              string                         `json:"connected_website,omitempty"`
    PassportData                  *PassportData                  `json:"passport_data,omitempty"`
    ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
    VoiceChatScheduled            *VoiceChatScheduled            `json:"voice_chat_scheduled,omitempty"`
    VoiceChatStarted              *VoiceChatStarted              `json:"voice_chat_started,omitempty"`
    VoiceChatEnded                *VoiceChatEnded                `json:"voice_chat_ended,omitempty"`
    VoiceChatParticipantsInvited  *VoiceChatParticipantsInvited  `json:"voice_chat_participants_invited,omitempty"`
    ReplyMarkup                   *InlineKeyboardReplyMarkup     `json:"reply_markup,omitempty"`
}

type Chat struct {
    Id                    int              `json:"id,omitempty"`
    Type                  string           `json:"type,omitempty"`
    Title                 string           `json:"title,omitempty"`
    Username              string           `json:"username,omitempty"`
    FirstName             string           `json:"first_name,omitempty"`
    LastName              string           `json:"last_name,omitempty"`
    Photo                 *ChatPhoto       `json:"photo,omitempty"`
    Bio                   string           `json:"bio,omitempty"`
    Description           string           `json:"description,omitempty"`
    InviteLink            string           `json:"invite_link,omitempty"`
    PinnedMessage         *Message         `json:"pinned_message,omitempty"`
    Permissions           *ChatPermissions `json:"permissions,omitempty"`
    SlowModeDelay         int              `json:"slow_mode_delay,omitempty"`
    MessageAutoDeleteTime int              `json:"message_auto_delete_time,omitempty"`
    StickerSetName        string           `json:"sticker_set_name,omitempty"`
    CanSetStickerSet      bool             `json:"can_set_sticker_set,omitempty"`
    LinkedChatId          int64            `json:"linked_chat_id,omitempty"`
    Location              *ChatLocation    `json:"location,omitempty"`
}

type ChatPhoto struct {
    SmallFileId       string `json:"small_file_id,omitempty"`
    SmallFileUniqueId string `json:"small_file_unique_id,omitempty"`
    BigFileId         string `json:"big_file_id,omitempty"`
    BigFileUniqueId   string `json:"big_file_unique_id,omitempty"`
}

type ChatPermissions struct {
    CanSendMessages       bool `json:"can_send_messages,omitempty"`
    CanSendMediaMessages  bool `json:"can_send_media_messages,omitempty"`
    CanSendPolls          bool `json:"can_send_polls,omitempty"`
    CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
    CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
    CanChangeInfo         bool `json:"can_change_info,omitempty"`
    CanInviteUsers        bool `json:"can_invite_users,omitempty"`
    CanPinMessages        bool `json:"can_pin_messages,omitempty"`
}
type ChatLocation struct {
    Location Location `json:"location,omitempty"`
    Address  string   `json:"address,omitempty"`
}

type Location struct {
    Longitude            float64 `json:"longitude,omitempty"`
    Latitude             float64 `json:"latitude,omitempty"`
    HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
    LivePeriod           int     `json:"live_period,omitempty"`
    Heading              int     `json:"heading,omitempty"`
    ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

type Animation struct {
    FileId       string     `json:"file_id,omitempty"`
    FileUniqueId string     `json:"file_unique_id,omitempty"`
    Width        int64      `json:"width,omitempty"`
    Height       int64      `json:"height,omitempty"`
    Duration     int64      `json:"duration,omitempty"`
    Thumb        *PhotoSize `json:"thumb,omitempty"`
    FileName     string     `json:"file_name,omitempty"`
    MimeType     string     `json:"mime_type,omitempty"`
    FileSize     int64      `json:"file_size,omitempty"`
}
type PhotoSize struct {
    FileId       string `json:"file_id,omitempty"`
    FileUniqueId string `json:"file_unique_id,omitempty"`
    Width        int    `json:"width,omitempty"`
    Height       int    `json:"height,omitempty"`
    FileSize     int    `json:"file_size,omitempty"`
}

type Audio struct {
    FileId       string     `json:"file_id,omitempty"`
    FileUniqueId string     `json:"file_unique_id,omitempty"`
    Duration     int64      `json:"duration,omitempty"`
    Performer    string     `json:"performer,omitempty"`
    Title        string     `json:"title,omitempty"`
    FileName     string     `json:"file_name,omitempty"`
    MimeType     string     `json:"mime_type,omitempty"`
    FileSize     int64      `json:"file_size,omitempty"`
    Thumb        *PhotoSize `json:"thumb,omitempty"`
}

type Document struct {
    FileId       string     `json:"file_id,omitempty"`
    FileUniqueId string     `json:"file_unique_id,omitempty"`
    Thumb        *PhotoSize `json:"thumb,omitempty"`
    FileName     string     `json:"file_name,omitempty"`
    MimeType     string     `json:"mime_type,omitempty"`
    FileSize     int64      `json:"file_size,omitempty"`
}

type Sticker struct {
    FileId       string        `json:"file_id,omitempty"`
    FileUniqueId string        `json:"file_unique_id,omitempty"`
    Width        int64         `json:"width,omitempty"`
    Height       int64         `json:"height,omitempty"`
    IsAnimated   bool          `json:"is_animated,omitempty"`
    Thumb        *PhotoSize    `json:"thumb,omitempty"`
    Emoji        string        `json:"emoji,omitempty"`
    SetName      string        `json:"set_name,omitempty"`
    MaskPosition *MaskPosition `json:"mask_position,omitempty"`
    FileSize     int64         `json:"file_size,omitempty"`
}

type MaskPosition struct {
    Point  string  `json:"point,omitempty"`
    XShift float64 `json:"x_shift,omitempty"`
    YShift float64 `json:"y_shift,omitempty"`
    Scale  float64 `json:"scale,omitempty"`
}

type Video struct {
    FileId       string     `json:"file_id,omitempty"`
    FileUniqueId string     `json:"file_unique_id,omitempty"`
    Width        int64      `json:"width,omitempty"`
    Height       int64      `json:"height,omitempty"`
    Duration     int64      `json:"duration,omitempty"`
    Thumb        *PhotoSize `json:"thumb,omitempty"`
    FileName     string     `json:"file_name,omitempty"`
    MimeType     string     `json:"mime_type,omitempty"`
    FileSize     int64      `json:"file_size,omitempty"`
}
type VideoNote struct {
    FileId       string     `json:"file_id,omitempty"`
    FileUniqueId string     `json:"file_unique_id,omitempty"`
    Length       int        `json:"length,omitempty"`
    Duration     int        `json:"duration,omitempty"`
    Thumb        *PhotoSize `json:"thumb,omitempty"`
    FileSize     int        `json:"file_size,omitempty"`
}
type Voice struct {
    FileId       string `json:"file_id,omitempty"`
    FileUniqueId string `json:"file_unique_id,omitempty"`
    Duration     int    `json:"duration,omitempty"`
    MimeType     string `json:"mime_type,omitempty"`
    FileSize     int    `json:"file_size,omitempty"`
}
type Contact struct {
    PhoneNumber string `json:"phone_number,omitempty"`
    FirstName   string `json:"first_name,omitempty"`
    LastName    string `json:"last_name,omitempty"`
    UserID      int64  `json:"user_id,omitempty"`
    VCard       string `json:"vcard,omitempty"`
}
type Dice struct {
    Emoji string `json:"emoji,omitempty"`
    Value int    `json:"value   ,omitempty"`
}

type Game struct {
    Title        string          `json:"title,omitempty"`
    Description  string          `json:"description,omitempty"`
    Photo        []PhotoSize     `json:"photo,omitempty"`
    Text         string          `json:"text,omitempty"`
    TextEntities []MessageEntity `json:"text_entities,omitempty"`
    Animation    Animation       `json:"animation,omitempty"`
}

type Poll struct {
    Id                    string          `json:"id,omitempty"`
    Question              string          `json:"question,omitempty"`
    Options               []PollOption    `json:"options,omitempty"`
    TotalVoterCount       int64           `json:"total_voter_count,omitempty"`
    IsClosed              bool            `json:"is_closed,omitempty"`
    IsAnonymous           bool            `json:"is_anonymous,omitempty"`
    Type                  string          `json:"type,omitempty"`
    AllowsMultipleAnswers bool            `json:"allows_multiple_answers,omitempty"`
    CorrectOptionId       int64           `json:"correct_option_id,omitempty"`
    Explanation           string          `json:"explanation,omitempty"`
    ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
    OpenPeriod            int64           `json:"open_period,omitempty"`
    CloseDate             int64           `json:"close_date,omitempty"`
}

type Venue struct {
    Location        Location `json:"location,omitempty"`
    Title           string   `json:"title,omitempty"`
    Address         string   `json:"address,omitempty"`
    FoursquareId    string   `json:"foursquare_id,omitempty"`
    FoursquareType  string   `json:"foursquare_type,omitempty"`
    GooglePlaceId   string   `json:"google_place_id,omitempty"`
    GooglePlaceType string   `json:"google_place_type,omitempty"`
}

type MessageAutoDeleteTimerChanged struct {
    MessageAutoDeleteTime int `json:"message_auto_delete_time,omitempty"`
}
type Invoice struct {
    Title          string `json:"title,omitempty"`
    Description    string `json:"description,omitempty"`
    StartParameter string `json:"start_parameter,omitempty"`
    Currency       string `json:"currency,omitempty"`
    TotalAmount    int    `json:"total_amount,omitempty"`
}
type SuccessfulPayment struct {
    Currency                string     `json:"currency,omitempty"`
    TotalAmount             int64      `json:"total_amount,omitempty"`
    InvoicePayload          string     `json:"invoice_payload,omitempty"`
    ShippingOptionId        string     `json:"shipping_option_id,omitempty"`
    OrderInfo               *OrderInfo `json:"order_info,omitempty"`
    TelegramPaymentChargeId string     `json:"telegram_payment_charge_id,omitempty"`
    ProviderPaymentChargeId string     `json:"provider_payment_charge_id,omitempty"`
}

type PassportData struct {
    Data        []EncryptedPassportElement `json:"data,omitempty"`
    Credentials EncryptedCredentials       `json:"credentials,omitempty"`
}

type InlineQuery struct {
    Id       string    `json:"id,omitempty"`
    From     User      `json:"from,omitempty"`
    Query    string    `json:"query,omitempty"`
    Offset   string    `json:"offset,omitempty"`
    ChatType string    `json:"chat_type,omitempty"`
    Location *Location `json:"location,omitempty"`
}
type ChosenInlineResult struct {
    ResultId        string    `json:"result_id,omitempty"`
    From            User      `json:"from,omitempty"`
    Location        *Location `json:"location,omitempty"`
    InlineMessageId string    `json:"inline_message_id,omitempty"`
    Query           string    `json:"query,omitempty"`
}

type ProximityAlertTriggered struct {
    Traveler User  `json:"traveler,omitempty"`
    Watcher  User  `json:"watcher,omitempty"`
    Distance int64 `json:"distance,omitempty"`
}

type VoiceChatEnded struct {
    Duration int64 `json:"duration,omitempty"`
}

type VoiceChatParticipantsInvited struct {
    Users []User `json:"users,omitempty"`
}

type VoiceChatScheduled struct {
    StartDate int64 `json:"start_date,omitempty"`
}

type VoiceChatStarted struct{}

type PollOption struct {
    Text       string `json:"text,omitempty"`
    VoterCount int    `json:"voter_count,omitempty"`
}
type OrderInfo struct {
    Name            string           `json:"name,omitempty"`
    PhoneNumber     string           `json:"phone_number,omitempty"`
    Email           string           `json:"email,omitempty"`
    ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type CallbackQuery struct {
    Id              string   `json:"id,omitempty"`
    From            User     `json:"from,omitempty"`
    Message         *Message `json:"message,omitempty"`
    InlineMessageId string   `json:"inline_message_id,omitempty"`
    ChatInstance    string   `json:"chat_instance,omitempty"`
    Data            string   `json:"data,omitempty"`
    GameShortName   string   `json:"game_short_name,omitempty"`
}
type PreCheckoutQuery struct {
    Id               string     `json:"id,omitempty"`
    From             User       `json:"from,omitempty"`
    Currency         string     `json:"currency,omitempty"`
    TotalAmount      int64      `json:"total_amount,omitempty"`
    InvoicePayload   string     `json:"invoice_payload,omitempty"`
    ShippingOptionId string     `json:"shipping_option_id,omitempty"`
    OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

type PollAnswer struct {
    PollId    string `json:"poll_id,omitempty"`
    User      User   `json:"user,omitempty"`
    OptionIDs []int  `json:"currency,omitempty"`
}

type ChatMemberUpdated struct {
    Chat          Chat            `json:"chat,omitempty"`
    From          User            `json:"from,omitempty"`
    Date          int             `json:"date,omitempty"`
    OldChatMember ChatMember      `json:"old_chat_member,omitempty"`
    NewChatMember ChatMember      `json:"new_chat_member,omitempty"`
    InviteLink    *ChatInviteLink `json:"invite_link,omitempty"`
}

// https://core.telegram.org/bots/api#chatmember
type ChatMember struct {
    User                  User   `json:"user,omitempty"`
    Status                string `json:"status,omitempty"`
    CustomTitle           string `json:"custom_title,omitempty"`
    IsAnonymous           bool   `json:"is_anonymous,omitempty"`
    CanBeEdited           bool   `json:"can_be_edited,omitempty"`
    CanManageChat         bool   `json:"can_manage_chat,omitempty"`
    CanPostMessages       bool   `json:"can_post_messages,omitempty"`
    CanEditMessages       bool   `json:"can_edit_messages,omitempty"`
    CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`
    CanManageVoiceChats   bool   `json:"can_manage_voice_chats,omitempty"`
    CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`
    CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`
    CanChangeInfo         bool   `json:"can_change_info,omitempty"`
    CanInviteUsers        bool   `json:"can_invite_users,omitempty"`
    CanPinMessages        bool   `json:"can_pin_messages,omitempty"`
    IsMember              bool   `json:"is_member,omitempty"`
    CanSendMessages       bool   `json:"can_send_messages,omitempty"`
    CanSendMediaMessages  bool   `json:"can_send_media_messages,omitempty"`
    CanSendPolls          bool   `json:"can_send_polls,omitempty"`
    CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`
    CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"`
    UntilDate             int64  `json:"until_date,omitempty"`
}

type ChatInviteLink struct {
    InviteLink  string `json:"invite_link,omitempty"`
    Creator     User   `json:"creator,omitempty"`
    IsPrimary   bool   `json:"is_primary,omitempty"`
    IsRevoked   bool   `json:"is_revoked,omitempty"`
    ExpireDate  int64  `json:"expire_date,omitempty"`
    MemberLimit int64  `json:"member_limit,omitempty"`
}

type ShippingAddress struct {
    CountryCode string `json:"country_code,omitempty"`
    State       string `json:"state,omitempty"`
    City        string `json:"city,omitempty"`
    StreetLine1 string `json:"street_line1,omitempty"`
    StreetLine2 string `json:"street_line2,omitempty"`
    PostCode    string `json:"post_code,omitempty"`
}

type ShippingQuery struct {
    Id              string          `json:"id,omitempty"`
    From            User            `json:"from,omitempty"`
    InvoicePayload  string          `json:"invoice_payload,omitempty"`
    ShippingAddress ShippingAddress `json:"shipping_address,omitempty"`
}

type EncryptedCredentials struct {
    Data   string `json:"data,omitempty"`
    Hash   string `json:"hash,omitempty"`
    Secret string `json:"secret,omitempty"`
}

type EncryptedPassportElement struct {
    Type        string         `json:"type,omitempty"`
    Data        string         `json:"data,omitempty"`
    PhoneNumber string         `json:"phone_number,omitempty"`
    Email       string         `json:"email,omitempty"`
    Files       []PassportFile `json:"files,omitempty"`
    FrontSide   *PassportFile  `json:"front_side,omitempty"`
    ReverseSide *PassportFile  `json:"reverse_side,omitempty"`
    Selfie      *PassportFile  `json:"selfie,omitempty"`
    Translation []PassportFile `json:"translation,omitempty"`
    Hash        string         `json:"hash,omitempty"`
}
type PassportFile struct {
    FileId       string `json:"file_id,omitempty"`
    FileUniqueId string `json:"file_unique_id,omitempty"`
    FileSize     int64  `json:"file_size,omitempty"`
    FileDate     int64  `json:"file_date,omitempty"`
}

///data

type MessageId struct {
    MessageId int `json:"message_id,omitempty"`
}

